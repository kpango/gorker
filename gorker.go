package gorker

import (
	"context"
	"sync"
)

type Dispatcher struct {
	running     bool
	queue       chan func()
	wg          *sync.WaitGroup
	workerCount int
	workers     []*worker
	ctx         context.Context
	cancel      context.CancelFunc
}

type worker struct {
	dis        *Dispatcher
	kill       chan struct{}
	processing bool
	running    bool
}

var (
	defaultWorker = 3
	instance      *Dispatcher
	once          sync.Once
)

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
	ctx, cancel := context.WithCancel(context.Background())
	dis := &Dispatcher{
		running:     false,
		workerCount: maxWorker,
		queue:       make(chan func(), maxWorker*1000),
		wg:          new(sync.WaitGroup),
		workers:     make([]*worker, maxWorker),
		ctx:         ctx,
		cancel:      cancel,
	}
	for i := 0; i < maxWorker; i++ {
		dis.workers[i] = &worker{
			dis:        dis,
			kill:       make(chan struct{}, 1),
			processing: false,
			running:    false,
		}
	}
	return dis
}

func (d *Dispatcher) StartWorkerObserver() *Dispatcher {
	go func() {
		for {
			select {
			case <-d.ctx.Done():
				return
			default:
				if len(d.workers) > d.workerCount {
					diff := len(d.workers) - d.workerCount
					idx := 0
					for {
						if !d.workers[idx].processing {
							if d.running && d.workers[idx].running {
								d.workers[idx].stop()
							}
							d.workers = append(d.workers[:idx], d.workers[idx:]...)
							diff--
							if diff == 0 {
								break
							}
						}
						idx++
						if idx == len(d.workers) {
							idx = 0
						}
					}
				} else if len(d.workers) < d.workerCount {
					diff := d.workerCount - len(d.workers)
					for {
						d.workers = append(d.workers, &worker{
							dis:        d,
							kill:       make(chan struct{}, 1),
							processing: false,
							running:    false,
						})
						diff--
						if diff == 0 {
							break
						}
					}
					for i, w := range d.workers {
						if d.running && !w.running {
							d.workers[i].start(d.ctx)
							d.workers[i].running = true
						}
					}
				}
			}
		}
	}()
	return d
}

func (d *Dispatcher) Reset(workerCount int) *Dispatcher {
	d.Stop(false)
	d = New(workerCount)
	return d
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

func (d *Dispatcher) Start() *Dispatcher {
	for _, worker := range d.workers {
		worker.start(d.ctx)
	}
	d.running = true
	return d
}

func (d *Dispatcher) Add(job func() error) chan error {
	d.wg.Add(1)
	ech := make(chan error, 1)
	d.queue <- func() {
		ech <- job()
	}
	return ech
}

func (d *Dispatcher) Wait() {
	if d.running {
		d.wg.Wait()
	}
}

func (d *Dispatcher) Stop(immediately bool) *Dispatcher {
	if !d.running {
		return d
	}

	if !immediately {
		d.wg.Wait()
	}

	d.cancel()
	close(d.queue)

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
				if job != nil {
					w.processing = true
					job()
					w.dis.wg.Done()
					w.processing = false
				}
			}
		}
	}()
}

func (w *worker) stop() {
	w.kill <- struct{}{}
	w.running = false
}
