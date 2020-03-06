package mq

import (
	"fmt"
	"net/http"
	"strconv"
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
				// go func(job Job) {
				// 	rdbc := Redis.Pool.Get() //获取一个redis连接
				// 	rdbc.Do("lpush", "zhengzijian", job.ID)
				// 	// fmt.Printf("ID: %d CT %v NOW %v\n", job.ID, job.CT, time.Now().Unix())
				// 	// fmt.Println(err)
				// 	rdbc.Close()
				// }(job)
				fmt.Println(job.ID)

				tr := &http.Transport{DisableKeepAlives: true}
				client := &http.Client{Transport: tr}

				resp, err := client.Get("http://192.168.6.21:3000/" + strconv.Itoa(job.ID))
				if resp != nil {
					defer resp.Body.Close()
				}

				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println(resp.StatusCode)
				// client := &http.Client{}
				// url := "http://192.168.6.21:3000/" + strconv.Itoa(job.ID)
				// reqest, err := http.NewRequest("GET", url, nil)
				// // defer reqest.Body.Close()
				// if err != nil {
				// 	panic(err)
				// }
				//处理返回结果
				// response, _ := client.Do(reqest)
				// reqest.Body.Close()
				//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
				// stdout := os.Stdout
				// _, err = io.Copy(stdout, response.Body)

				// //返回的状态码
				// status := response.StatusCode
				fmt.Println(job.ID, resp.StatusCode)

				// fmt.Println(status)

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
