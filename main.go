package main

import (
	"ymq/mq"
)

func main() {
	mq.JobQueueNew()
	runMQ := mq.NewYMQ(5)
	runMQ.Run()
}
