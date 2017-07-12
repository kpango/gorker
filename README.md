# gorker [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![release](https://img.shields.io/github/release/kpango/gorker.svg)](https://github.com/kpango/gorker/releases/latest) [![CircleCI](https://circleci.com/gh/kpango/gorker.svg?style=shield)](https://circleci.com/gh/kpango/gorker) [![codecov](https://codecov.io/gh/kpango/gorker/branch/master/graph/badge.svg)](https://codecov.io/gh/kpango/gorker) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/a6e544eee7bc49e08a000bb10ba3deed)](https://www.codacy.com/app/i.can.feel.gravity/gorker?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kpango/gorker&amp;utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/github.com/kpango/gorker)](https://goreportcard.com/report/github.com/kpango/gorker) [![GoDoc](http://godoc.org/github.com/kpango/gorker?status.svg)](http://godoc.org/github.com/kpango/gorker) [![Join the chat at https://gitter.im/kpango/gorker](https://badges.gitter.im/kpango/gorker.svg)](https://gitter.im/kpango/gorker?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

gorker is golang dispatch worker management library

## Requirement
Go 1.8

## Installation
```shell
go get github.com/kpango/gorker
```

## Example
```go
func main() {
	dispatcher := gorker.Get(3)
	dispatcher.StartWorkerObserver()

	for i := 0; i < 10000; i++ {
		func(n int) {
			dispatcher.Add(func() error {
				fmt.Printf("%03d:\t workers: %d\t%v\n", n, runtime.NumGoroutine()-2, time.Now().Format(time.RFC3339))
				time.Sleep(time.Millisecond * 100)
				return nil
			})
		}(i)
	}
	dispatcher.Start()

	time.Sleep(time.Second * 5)

	gorker.UpScale(7)
	fmt.Printf("UpScale : %d\n", 7)

	time.Sleep(time.Second * 5)

	gorker.DownScale(2)
	fmt.Printf("DownScale : %d\n", 2)

	time.Sleep(time.Second * 5)

	gorker.UpScale(20)
	time.Sleep(time.Second * 5)

	dispatcher.Add(func() error {
		fmt.Printf("last worker:\t workers: %d\t%v\n", runtime.NumGoroutine()-2, time.Now().Format(time.RFC3339))
		time.Sleep(time.Millisecond * 100)
		return nil
	})

	gorker.UpScale(200)
	time.Sleep(time.Second * 5)

	dispatcher.Stop(true)

	dispatcher.Wait()
}
```

![Sample Logs](https://github.com/kpango/gorker/raw/master/images/sample.png)

## Benchmarks

![Bench](https://github.com/kpango/gorker/raw/master/images/bench.png)

## Contribution
1. Fork it ( https://github.com/kpango/gorker/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## Author
[kpango](https://github.com/kpango)

## LICENSE
gorker released under MIT license, refer [LICENSE](https://github.com/kpango/gorker/blob/master/LICENSE) file.
