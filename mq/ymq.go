package mq

import (
	"context"
	"fmt"
)

var (
	ymq *YMQ
)

//YMQ YMessgaeQueue
type YMQ struct {
	status     int
	closed     chan struct{}
	dispatcher *Dispatcher
	rpc        *RPC
}

//NewYMQ YMQ工厂
func NewYMQ(maxWorkers int) *YMQ {
	ymq = &YMQ{
		closed: make(chan struct{}),
	}
	fmt.Println("创建YMQ")
	ymq.dispatcher = NewDispatcher(maxWorkers)
	ymq.rpc = NewRPC()

	return ymq
}

//Run 运行YMQ
func (ymq *YMQ) Run() {
	if ymq.status == 1 {
		fmt.Println("ymq is running now.")
		return
	}

	ctx, cannel := context.WithCancel(context.Background())
	defer cannel()

	ymq.status = 1

	go ymq.dispatcher.Run()
	go ymq.rpc.Run(ctx)

	<-ymq.closed
	fmt.Println("Closed.")
}
