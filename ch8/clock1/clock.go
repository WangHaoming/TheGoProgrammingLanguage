package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().UTC().Format("15:04:05\n"))
		if err != nil {
			return
		}
		// 这个函数会阻塞，直到遇到一个EOF才返回（ctrl+D）
		io.Copy(os.Stdout, c)
		fmt.Println("got message!")
		time.Sleep(time.Second)
	}
}
