package main

import (
	"GoSpider/ConcurrenceTask/learngo/rpc"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":8088")

	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	args := rpc.Args{A: 14, B: 7}
	var result float64

	err = client.Call("DemoService.Div", args, &result)

	if err != nil {
		log.Printf("error:%v\n", err)
	} else {
		log.Printf("%d / %d = %.5f\n", args.A, args.B, result)
	}

	args = rpc.Args{A: 14, B: 0}

	err = client.Call("DemoService.Div", args, &result)

	if err != nil {
		log.Printf("error:%v\n", err)
	} else {
		log.Printf("%d / %d = %.5f\n", args.A, args.B, result)
	}


}
