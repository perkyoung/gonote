package main

import (
	"fmt"
	"sync"
)

type Worker2 interface {
	Task()
}

type Myworker2 struct {
	name string
}

func (my *Myworker2) Task() {
	fmt.Println("my name is ", my.name)
}

type WorkerPool struct {
	workerChan chan Worker2
	wg         sync.WaitGroup
}

func NewWorkerPool(gonum int) *WorkerPool{
	workpool := WorkerPool{workerChan: make(chan Worker2)}
	workpool.wg.Add(gonum)
	for i := 0; i <= gonum; i++ {
		go func() {
			for msg := range workpool.workerChan {
				msg.Task()
			}
		}()
		workpool.wg.Done()
	}
	return &workpool
}