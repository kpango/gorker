package main

import (
	"runtime"
	"time"

	"github.com/kpango/glg"
	"github.com/kpango/gorker"
)

func main() {
	dispatcher := gorker.Get(3).StartWorkerObserver().QueueRunner()

	go func() {
		for i := 0; i < 10000000; i++ {
			// for i := 0; i < 100000; i++ {
			func(n int) {
				dispatcher.Add(func() error {
					glg.Infof("%03d:\t workers: %d\n", n, runtime.NumGoroutine()-2)
					// time.Sleep(time.Millisecond * 10)
					time.Sleep(time.Millisecond * 100)
					return nil
				})
			}(i)
		}
		glg.Debug("done")
	}()

	dispatcher.Start()

	time.Sleep(time.Second)

	gorker.UpScale(7)
	glg.Debugf("UpScale : %d\n", 7)

	time.Sleep(time.Second)

	gorker.DownScale(2)
	glg.Debugf("DownScale : %d\n", 2)

	time.Sleep(time.Second)

	gorker.UpScale(20)
	time.Sleep(time.Second)

	dispatcher.Add(func() error {
		glg.Successf("last worker:\t workers: %d\n", runtime.NumGoroutine()-2)
		time.Sleep(time.Millisecond * 100)
		return nil
	})

	gorker.UpScale(200)

	time.Sleep(10 * time.Second)

	gorker.UpScale(2000)

	dispatcher.Stop(false)
}
