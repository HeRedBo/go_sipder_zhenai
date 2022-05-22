package main

import (
	rpc2 "GoSpider/ConcurrenceTask/learngo/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


func main() {
	rpc.Register(rpc2.DemoService{})
	listener, err := net.Listen("tcp",":8088")
	if err != nil {
		panic(err)
	}

	for {
		conn ,err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
