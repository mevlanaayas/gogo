
package main

import (
	"net/rpc"
	"net"
	"fmt"
)

type Server struct{}

func (server *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func main()  {
	go rpcServer()
	go rpcClient()

	var input string
	fmt.Scanln(&input)

}

func rpcServer()  {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)

	}

}

func rpcClient()  {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) = ", result)
	}
	c.Close()

}