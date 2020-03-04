package mq

//Dispatcher 创建worker池
type Dispatcher struct {
	WorkerPool chan chan Job
}

//NewDispatcher Dispatcher工厂
func NewDispatcher(maxWokers int) *Dispatcher {
	pool := make(chan chan Job, maxWokers)
	return &Dispatcher{WorkerPool: pool}
}

//dispatch
func (dispatcher *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue.Jobs:
			go func(job Job) {
				jobChannel := <-dispatcher.WorkerPool

				jobChannel <- job
			}(job)
		}
	}
}

//Run dispacther启动
func (dispatcher *Dispatcher) Run() {
	// println("worker启动")
	// println(dispatcher.maxWokers)
	maxWorkers := 5

	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(dispatcher.WorkerPool)
		worker.Start()
	}

	go dispatcher.dispatch()
}
