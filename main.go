package main

import (
	"ymq/mq"
)

func main() {
	mq.JobQueueNew(40000)
	runMQ := mq.NewYMQ(10)
	runMQ.Run()
}
