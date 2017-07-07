package gorker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kpango/gorker"
)

func Test_GospatchWorker(t *testing.T) {
	dispatcher := gorker.Get(3)
	dispatcher.StartWorkerObserver()

	for i := 0; i < 100; i++ {
		func(n int) {
			dispatcher.Add(func() error {
				fmt.Printf("%03d: aaaaaa\n", n)
				time.Sleep(time.Second * 2)
				return nil
			})
		}(i)
	}
	dispatcher.Start()

	time.Sleep(time.Second * 10)

	gorker.Get(20)

	dispatcher.Wait()
}
