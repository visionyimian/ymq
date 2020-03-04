package mq

import (
	"fmt"
	"time"
)

//Worker 通过Worker具体执行具体的Job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

//Start Worker启动
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				//do something
				// println(job.ID, time.Now())
				fmt.Printf("ID: %d CT %v NOW %v\n", job.ID, job.CT, time.Now().Unix())
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop Worker停止
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//NewWorker Worker工厂
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}
