package main

import (
	"encoding/json"
	"fmt"
	"github.com/spiral/goridge"
	"net"
	"net/rpc"
	"strconv"
)

type App struct{}

func (s *App) SumArray(array string, r *string) error {
	var arrayStruct []int
	err := json.Unmarshal([]byte(array), &arrayStruct)
	if err != nil {
		return err
	}
	sum := 0
	returnString := ""
	for i, value := range arrayStruct {
		sum += value
		returnString += strconv.Itoa(value)
		if i != len(arrayStruct)-1 {
			returnString += "+"
		}
	}
	*r = fmt.Sprintf(returnString+"=%d\n", sum)
	return nil
}

func main() {
	ln, err := net.Listen("tcp", ":6001")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen tcp:6001")
	err = rpc.Register(new(App))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		fmt.Println("accept connection from", conn.RemoteAddr().String())
		go rpc.ServeCodec(goridge.NewCodec(conn))
	}
}
