package gorker

import (
	"reflect"
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
			got.Start()
			if !got.running {
				t.Error("worker is not running")
			}
			if got.workerCount != tt.maxWorker {
				t.Errorf("workerCount = %v, want %v", got.workerCount, tt.maxWorker)
			}
			if len(got.workers) != tt.maxWorker {
				t.Errorf("worker length = %v, want %v", len(got.workers), tt.maxWorker)
			}
			got.Stop(false)
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
	ins1 := newDispatcher(defaultWorker)
	ins2 := newDispatcher(defaultWorker)

	if ins1.workerCount != ins2.workerCount {
		t.Errorf("got = %v, want = %v", ins1.workerCount, ins2.workerCount)
	}

	if len(ins1.workers) != len(ins2.workers) {
		t.Errorf("worker length is %v, wnat length = %v", len(ins1.workers), len(ins2.workers))
	}
}

func TestDispatcher_GetWorkerCount(t *testing.T) {
	got := Get(10)

	if got.workerCount != got.GetWorkerCount() {
		t.Error(got.GetWorkerCount())
		t.Error("invalid worker count")
	}
}
