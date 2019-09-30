package main

import (
	"fmt"
	"net"

	"github.com/recoilme/mcproto"
)

func main() {
	listener, err := net.Listen("tcp", ":11212")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// start
	kuku := NewKuku()
	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("conn", err)
			conn.Close()
			continue
		}
		go mcproto.ParseMc(conn, kuku, "") //listen(conn, db)
	}
}
