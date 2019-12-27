package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	lister,err := net.Listen("tcp", host)
	if err != nil{
		return err
	}

	fmt.Printf("Listening on %s", host)
	for {
		conn, err := lister.Accept()
		if err !=nil {
			log.Printf("conn errors: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn,err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(conn)
	return client, nil
}