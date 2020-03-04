package mq

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

//Service 结构体
type Service struct {
}

//Push 推送信息
func (service *Service) Push(ID int, reply *string) error {
	jobQueue.Jobs <- Job{ID: ID, CT: time.Now().Unix()}
	*reply = "success"
	return nil
}

//RPC 结构体
type RPC struct {
}

//NewRPC 创建PPC
func NewRPC() *RPC {
	rpc := &RPC{}
	return rpc
}

//Run RPC服务端启动
func (r *RPC) Run(ctx context.Context) {
	rpc.Register(new(Service))
	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		fmt.Println("rpc listen port error")
	} else {
		defer listener.Close()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
		go jsonrpc.ServeConn(conn)
	}
}
