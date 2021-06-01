package main

import (
	"fmt"
	"sync"
)

//场景，100个任务，分给n个协程去完成

type Worker interface {
	Task()
}
type MyWorker struct {
	name string
}

func (w *MyWorker) Task() {
	fmt.Println("my name is ", w.name)
}

type workPool struct {
	workers chan Worker
	wg      sync.WaitGroup
}

func NewWorkPool(gonum int) *workPool {
	p := workPool{workers: make(chan Worker)}
	p.wg.Add(gonum)
	for i := 0; i < gonum; i++ {
		go func() {
			for w := range p.workers {	//结束的时候，一定要先调用close(p.workers)，才会跳出for，到达后面的done，否则wg.wait永远无法结束
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *workPool) Run(w Worker) {
	p.workers <- w
}

func (p *workPool) Shutdown() {
	close(p.workers)	//注意，这里，一定要先close , 上面的逻辑才会调用到wg.done()，否则会永远等待
	p.wg.Wait()
}

var names = []string{
	"name",
	"perkoung",
	"liguoli",
	"marg",
	"therese",
}

func main() {
	p := NewWorkPool(3)
	var wg sync.WaitGroup

	wg.Add(10 * len(names))
	for i := 0; i < 10; i++ {
		for _, name := range names {
			np := MyWorker{name: name}
			go func() {
				p.Run(&np)
				wg.Done()
			}()

		}
	}
	wg.Wait()
	p.Shutdown()
}
