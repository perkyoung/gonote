package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

//指定超时时间，运行n个任务，捕获信号

var errInterrupt = errors.New("signal interrupt error")
var errTimeout = errors.New("timeout error")

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	task      []func(int)
}

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error, 1),
		timeout:   time.After(d),
	}
}

func (r *Runner) getInterrupt() error {
	select {
	case <-r.interrupt:
		return errInterrupt
	default:
		return nil
	}
}

func (r *Runner) run() error {
	for index, task := range r.task {
		if err := r.getInterrupt(); err != nil {
			return err
		}
		task(index)
	}
	return nil
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt, os.Kill)

	go func() {
		r.complete <- r.run()
	}()

	// 等待完成，或者等待超时
	select {
	case err := <-r.complete:	//complete 通道包含捕获信号提前结束
		return err
	case <-r.timeout:
		return errTimeout
	}
}

func (r *Runner) Add(task ...func(int)) {
	r.task = append(r.task, task...)
}

func testRunner1(i int) {
	fmt.Println("hello, test1 ", i)
}
func testRunner2(i int) {
	fmt.Println("hello, test2 ", i)
}

func main() {
	r := New(time.Second * 10)
	r.Add(testRunner1, testRunner2)
	r.Start()
}
