package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/kpango/gorker"
)

func main() {
	dispatcher := gorker.Get(3).StartWorkerObserver().QueueRunner()

	go func() {
		for i := 0; i < 10000000; i++ {
			func(n int) {
				dispatcher.Add(func() error {
					fmt.Printf("%03d:\t workers: %d\t%v\n", n, runtime.NumGoroutine()-2, time.Now().Format(time.RFC3339))
					time.Sleep(time.Millisecond * 10)
					return nil
				})
			}(i)
		}
	}()

	dispatcher.Start()

	time.Sleep(time.Second)

	gorker.UpScale(7)
	fmt.Printf("UpScale : %d\n", 7)

	time.Sleep(time.Second)

	gorker.DownScale(2)
	fmt.Printf("DownScale : %d\n", 2)

	time.Sleep(time.Second)

	gorker.UpScale(20)
	time.Sleep(time.Second)

	dispatcher.Add(func() error {
		fmt.Printf("last worker:\t workers: %d\t%v\n", runtime.NumGoroutine()-2, time.Now().Format(time.RFC3339))
		time.Sleep(time.Millisecond * 100)
		return nil
	})

	gorker.UpScale(200)

	time.Sleep(10 * time.Second)

	gorker.UpScale(2000)

	dispatcher.Stop(false)
}
