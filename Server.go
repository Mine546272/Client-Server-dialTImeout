package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3355")
	if err != nil {
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		connection(conn)
	}

}
func connection(conn net.Conn) {
	readDeadline := time.Duration(5 * time.Second)
	deadlineread := time.Now().Add(readDeadline)
	err := conn.SetReadDeadline(deadlineread)
	if err != nil {
		panic(err)
	}
	var size uint32
	err = binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	trfmsg := make([]byte, size)
	_, err = conn.Read(trfmsg)
	if err != nil {
		panic(err)
	}

	stringmsg := string(trfmsg)
	fmt.Printf("Message : %s", stringmsg)

	readDeadline2 := time.Duration(5 * time.Second)
	deadlineread2 := time.Now().Add(readDeadline2)
	err = conn.SetReadDeadline(deadlineread2)
	if err != nil {
		panic(err)
	}
}
