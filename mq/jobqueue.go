package mq

var (
	jobQueue *JobQueue
)

//JobQueue 结构体
type JobQueue struct {
	Jobs chan Job
}

//JobQueueNew 工厂方法
func JobQueueNew(q int64) *JobQueue {
	jobQueue = &JobQueue{
		Jobs: make(chan Job, q)}
	return jobQueue
}
