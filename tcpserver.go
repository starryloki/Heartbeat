package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func StartServer(port string) {
	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Printf("listen failed, err:%v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept failed, err:%v\n", err)
			continue
		}
		err1 := conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		if err1 != nil {
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			//log.Printf("read from conn failed, err:%v\n", err)
			break
		}

		//recv := string(buf[:n])

		_, err = conn.Write([]byte("ok"))
		if err != nil {
			log.Printf("write from conn failed, err:%v\n", err)
			break
		}
		contextHandler(buf[:n])
	}
}
