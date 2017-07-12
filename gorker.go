package gorker

import (
	"context"
	"sync"
)

type Dispatcher struct {
	running     bool
	queue       chan func()
	wg          *sync.WaitGroup
	mu          *sync.Mutex
	workerCount int
	workers     []*worker
	ctx         context.Context
	cancel      context.CancelFunc
}

type worker struct {
	dis     *Dispatcher
	kill    chan struct{}
	running bool
}

var (
	defaultWorker = 3
	instance      *Dispatcher
	once          sync.Once
)

func init() {
	instance = New(defaultWorker)
}

func GetInstance() *Dispatcher {
	return Get(defaultWorker)
}

func Get(maxWorker int) *Dispatcher {
	once.Do(func() {
		instance = New(maxWorker)
	})
	instance.workerCount = maxWorker
	return instance
}

func New(maxWorker int) *Dispatcher {
	if maxWorker < 1 {
		maxWorker = 1
	}
	dis := newDispatcher(maxWorker)
	for i := range dis.workers {
		dis.workers[i] = &worker{
			dis:     dis,
			kill:    make(chan struct{}, 1),
			running: false,
		}
	}
	return dis
}

func newDispatcher(maxWorker int) *Dispatcher {
	ctx, cancel := context.WithCancel(context.Background())
	return &Dispatcher{
		running:     false,
		workerCount: maxWorker,
		queue:       make(chan func(), 100000),
		wg:          new(sync.WaitGroup),
		mu:          new(sync.Mutex),
		workers:     make([]*worker, maxWorker),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (d *Dispatcher) SetQueueSize(size int) *Dispatcher {
	old := d.queue
	d.queue = make(chan func(), size)
	go func() {
		for job := range old {
			d.queue <- job
		}
	}()
	return d
}

func UpScale(workerCount int) *Dispatcher {
	return instance.UpScale(workerCount)
}

func (d *Dispatcher) UpScale(workerCount int) *Dispatcher {
	d.mu.Lock()
	diff := workerCount - len(d.workers)
	for {
		if diff < 1 {
			break
		}
		d.workers = append(d.workers, &worker{
			dis:     d,
			kill:    make(chan struct{}, 1),
			running: false,
		})
		diff--
	}
	d.workerCount = workerCount
	d.mu.Unlock()
	if d.running {
		d.Start()
	}
	return d
}

func DownScale(workerCount int) *Dispatcher {
	return instance.DownScale(workerCount)
}

func (d *Dispatcher) DownScale(workerCount int) *Dispatcher {
	d.mu.Lock()
	diff := len(d.workers) - workerCount
	idx := 0
	for {
		if diff < 1 {
			break
		}
		if d.running && d.workers[idx].running {
			d.workers[idx].stop()
		}
		d.workers = append(d.workers[:idx], d.workers[idx+1:]...)
		diff--
		idx++
		if idx >= len(d.workers) {
			idx = 0
		}
	}
	d.workerCount = workerCount
	d.mu.Unlock()
	return d
}

func AutoScale() *Dispatcher {
	return instance.AutoScale()
}

func (d *Dispatcher) AutoScale() *Dispatcher {
	d.mu.Lock()
	if len(d.workers) > d.workerCount {
		d.mu.Unlock()
		d.DownScale(d.workerCount)
	} else if len(d.workers) < d.workerCount {
		d.mu.Unlock()
		d.UpScale(d.workerCount)
	} else {
		d.mu.Unlock()
	}
	return d
}

func StartWorkerObserver() *Dispatcher {
	return instance.StartWorkerObserver()
}

func (d *Dispatcher) StartWorkerObserver() *Dispatcher {
	go func() {
		for {
			select {
			case <-d.ctx.Done():
				return
			default:
				if d.workerCount != len(d.workers) {
					d.AutoScale()
				}
			}
		}
	}()
	return d
}

func Reset(workerCount int) *Dispatcher {
	return instance.Reset(workerCount)
}

func (d *Dispatcher) Reset(workerCount int) *Dispatcher {
	d.Stop(false)
	d = New(workerCount)
	return d
}

func StartWithContext(c context.Context) *Dispatcher {
	return instance.StartWithContext(c)
}

func (d *Dispatcher) StartWithContext(c context.Context) *Dispatcher {
	ctx, cancel := context.WithCancel(c)
	d.ctx = ctx
	d.cancel = cancel
	for _, worker := range d.workers {
		worker.start(d.ctx)
	}
	d.running = true
	return d
}

func Start() *Dispatcher {
	return instance.Start()
}

func (d *Dispatcher) Start() *Dispatcher {
	for i, w := range d.workers {
		if !w.running {
			d.workers[i].start(d.ctx)
			d.workers[i].running = true
		}
	}
	d.running = true
	return d
}

func Add(job func() error) chan error {
	return instance.Add(job)
}

func (d *Dispatcher) Add(job func() error) chan error {
	d.wg.Add(1)
	ech := make(chan error, 1)
	d.queue <- func() {
		ech <- job()
	}
	return ech
}

func Wait() {
	instance.Wait()
}

func (d *Dispatcher) Wait() {
	if d.running {
		d.wg.Wait()
	}
}

func Stop(immediately bool) *Dispatcher {
	return instance.Stop(immediately)
}

func (d *Dispatcher) Stop(immediately bool) *Dispatcher {
	if !d.running {
		return d
	}

	if !immediately {
		d.Wait()
	}

	d.cancel()

	d.running = false
	d = New(len(d.workers))
	return d
}

func (w *worker) start(ctx context.Context) {
	w.running = true
	go func() {
		for {
			select {
			case <-w.kill:
				return
			case <-ctx.Done():
				return
			case job := <-w.dis.queue:
				w.run(job)
			}
		}
	}()
}

func (w *worker) run(job func()) {
	if job != nil {
		job()
		w.dis.wg.Done()
	}
}

func (w *worker) stop() {
	w.kill <- struct{}{}
	w.running = false
}
