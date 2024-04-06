package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	dial, err := net.DialTimeout("tcp", "localhost:3355", 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	writeDeadline := time.Duration(5 * time.Second)
	deadlinewrite := time.Now().Add(writeDeadline)
	err = dial.SetWriteDeadline(deadlinewrite)
	if err != nil {
		panic(err)
	}
	message := "This is The message......"
	err = binary.Write(dial, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}

	_, err = dial.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
}
