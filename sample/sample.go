package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/kpango/gorker"
)

func main() {
	d := gorker.Get(10).StartWorkerObserver().QueueRunner()

	go func() {
		for i := 0; i < 10000000; i++ {
			func(n int) {
				d.Add(func() error {
					fmt.Printf("%03d:\t workers: %d\t%v\n", n, runtime.NumGoroutine()-2, time.Now().Format(time.RFC3339))
					time.Sleep(time.Millisecond * 10)
					return nil
				})
			}(i)
		}
	}()

	d.Start()

	time.Sleep(time.Second)

	gorker.UpScale(100)

	time.Sleep(time.Second * 5)

	gorker.DownScale(2)

	gorker.Wait()
}
