package mq

var (
	jobQueue *JobQueue
)

//JobQueue 结构体
type JobQueue struct {
	Jobs chan Job
}

//JobQueueNew 工厂方法
func JobQueueNew() *JobQueue {
	jobQueue = &JobQueue{
		Jobs: make(chan Job)}
	return jobQueue
}
