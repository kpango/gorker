package gorker

import (
	"context"
	"reflect"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		want *Dispatcher
	}{
		{
			name: "GetInstance equality",
			want: GetInstance(),
		},
		{
			name: "Get defaultWorker equality",
			want: Get(defaultWorker),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	workerPattern := []int{
		-1,
		0,
		1,
		10,
		100,
	}
	tests := []struct {
		name      string
		want      *Dispatcher
		maxWorker int
	}{
		{
			name:      "-1 pattern",
			want:      Get(workerPattern[0]),
			maxWorker: 1,
		},
		{
			name:      "0 pattern",
			want:      Get(workerPattern[1]),
			maxWorker: 1,
		},
		{
			name:      "1 pattern",
			want:      Get(workerPattern[2]),
			maxWorker: workerPattern[2],
		},
		{
			name:      "10 pattern",
			want:      Get(workerPattern[3]),
			maxWorker: workerPattern[3],
		},
		{
			name:      "100 pattern",
			want:      Get(workerPattern[4]),
			maxWorker: workerPattern[4],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Get(tt.maxWorker)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
			if got.workerCount != tt.maxWorker {
				t.Errorf("workerCount = %v, want %v", got.workerCount, tt.maxWorker)
			}
			if len(got.workers) != tt.maxWorker {
				t.Errorf("worker length = %v, want %v", len(got.workers), tt.maxWorker)
			}
			if got.running {
				t.Error("worker is running")
			}
			got.Start()
			if !got.running {
				t.Error("worker is not running")
			}
			got.Stop(true)
			if got.running {
				t.Error("worker is running")
			}
		})
	}
}

func TestNew(t *testing.T) {
	workerPattern := []int{
		-1,
		0,
		1,
		10,
		100,
	}
	tests := []struct {
		name      string
		want      *Dispatcher
		maxWorker int
	}{
		{
			name:      "-1 pattern",
			want:      New(workerPattern[0]),
			maxWorker: 1,
		},
		{
			name:      "0 pattern",
			want:      New(workerPattern[1]),
			maxWorker: 1,
		},
		{
			name:      "1 pattern",
			want:      New(workerPattern[2]),
			maxWorker: workerPattern[2],
		},
		{
			name:      "10 pattern",
			want:      New(workerPattern[3]),
			maxWorker: workerPattern[3],
		},
		{
			name:      "100 pattern",
			want:      New(workerPattern[4]),
			maxWorker: workerPattern[4],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.maxWorker)
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want different %v", got, tt.want)
			}
			if got.workerCount != tt.maxWorker {
				t.Errorf("workerCount = %v, want %v", got.workerCount, tt.maxWorker)
			}
			if len(got.workers) != tt.maxWorker {
				t.Errorf("worker length = %v, want %v", len(got.workers), tt.maxWorker)
			}
			if got.running {
				t.Error("worker is running")
			}
			got.Start()
			if !got.running {
				t.Error("worker is not running")
			}
			got.Stop(true)
			if got.running {
				t.Error("worker is running")
			}
		})
	}
}

func Test_newDispatcher(t *testing.T) {
	type args struct {
		maxWorker int
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDispatcher(tt.args.maxWorker); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_GetCurrentWorkerCount(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.GetCurrentWorkerCount(); got != tt.want {
				t.Errorf("Dispatcher.GetCurrentWorkerCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_SetQueueSize(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.SetQueueSize(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.SetQueueSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpScale(t *testing.T) {
	type args struct {
		workerCount int
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpScale(tt.args.workerCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_UpScale(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	type args struct {
		workerCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.UpScale(tt.args.workerCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.UpScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownScale(t *testing.T) {
	type args struct {
		workerCount int
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DownScale(tt.args.workerCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_DownScale(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	type args struct {
		workerCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.DownScale(tt.args.workerCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.DownScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoScale(t *testing.T) {
	tests := []struct {
		name string
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AutoScale(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_AutoScale(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.AutoScale(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.AutoScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartWorkerObserver(t *testing.T) {
	tests := []struct {
		name string
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartWorkerObserver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartWorkerObserver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_StartWorkerObserver(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.StartWorkerObserver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.StartWorkerObserver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReset(t *testing.T) {
	tests := []struct {
		name string
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_Reset(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.Reset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeReset(t *testing.T) {
	tests := []struct {
		name string
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeReset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeReset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_SafeReset(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.SafeReset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.SafeReset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartWithContext(t *testing.T) {
	type args struct {
		c context.Context
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartWithContext(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_StartWithContext(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	type args struct {
		c context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.StartWithContext(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.StartWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name string
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Start(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_Start(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.Start(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		job func() error
	}
	tests := []struct {
		name string
		args args
		want chan error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.job); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_Add(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	type args struct {
		job func() error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   chan error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.Add(tt.args.job); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWait(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Wait()
		})
	}
}

func TestDispatcher_Wait(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			d.Wait()
		})
	}
}

func TestStop(t *testing.T) {
	type args struct {
		immediately bool
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stop(tt.args.immediately); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_Stop(t *testing.T) {
	type fields struct {
		running     bool
		scaling     bool
		queue       chan func()
		wg          *sync.WaitGroup
		mu          *sync.Mutex
		workerCount int
		workers     []*worker
		ctx         context.Context
		cancel      context.CancelFunc
	}
	type args struct {
		immediately bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Dispatcher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				running:     tt.fields.running,
				scaling:     tt.fields.scaling,
				queue:       tt.fields.queue,
				wg:          tt.fields.wg,
				mu:          tt.fields.mu,
				workerCount: tt.fields.workerCount,
				workers:     tt.fields.workers,
				ctx:         tt.fields.ctx,
				cancel:      tt.fields.cancel,
			}
			if got := d.Stop(tt.args.immediately); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dispatcher.Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worker_start(t *testing.T) {
	type fields struct {
		dis     *Dispatcher
		kill    chan struct{}
		running bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &worker{
				dis:     tt.fields.dis,
				kill:    tt.fields.kill,
				running: tt.fields.running,
			}
			w.start(tt.args.ctx)
		})
	}
}

func Test_worker_run(t *testing.T) {
	type fields struct {
		dis     *Dispatcher
		kill    chan struct{}
		running bool
	}
	type args struct {
		job func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &worker{
				dis:     tt.fields.dis,
				kill:    tt.fields.kill,
				running: tt.fields.running,
			}
			w.run(tt.args.job)
		})
	}
}

func Test_worker_stop(t *testing.T) {
	type fields struct {
		dis     *Dispatcher
		kill    chan struct{}
		running bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &worker{
				dis:     tt.fields.dis,
				kill:    tt.fields.kill,
				running: tt.fields.running,
			}
			w.stop()
		})
	}
}
