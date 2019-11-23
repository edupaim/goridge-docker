package main

import (
	"fmt"
	"github.com/spiral/goridge"
	"net"
	"net/rpc"
)

type App struct{}

func (s *App) Hi(name string, r *string) error {
	*r = fmt.Sprintf("Hello, %s!", name)
	return nil
}

func main() {
	ln, err := net.Listen("tcp", ":6001")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen tcp:6001")
	rpc.Register(new(App))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		fmt.Println("accept connection from", conn.RemoteAddr().String())
		go rpc.ServeCodec(goridge.NewCodec(conn))
	}
}
