package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp",host)
	if err != nil {
		return err
	}
	log.Printf("Listen on %s", host)

	for {
		conn ,err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}


func NewCient(host string) (*rpc.Client,error) {
	conn , err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}

