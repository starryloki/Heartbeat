package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var user []string
var pipe chan string
var serverMode bool
var clientMode bool
var clientIp string
var clientName string
var serverPort string

func main() {
	flag.BoolVar(&serverMode, "s", false, "server")
	flag.BoolVar(&clientMode, "c", false, "client")
	flag.StringVar(&clientIp, "i", "", "ip:port")
	flag.StringVar(&clientName, "n", "", "client name")
	flag.StringVar(&serverPort, "p", "9090", "server port")
	flag.Parse()
	if serverMode == true {
		pipe = make(chan string)
		StartServer(serverPort)
	}
	if clientMode == true {
		for {
			sent(clientIp, clientName)
			time.Sleep(10 * time.Second)
		}
	}
	log.Print("please select a mode to run")
}

func sent(ip string, contextS string) {
	contextS = "HEARTB%" + contextS
	contextS = strings.TrimSpace(contextS)
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Printf("conn server failed, err:%v\n", err)
		return
	}
	_, err = conn.Write([]byte(contextS))
	if err != nil {
		log.Printf("send failed, err:%v\n", err)
		return
	}
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		log.Printf("read failed:%v\n", err)
		return
	}
	fmt.Printf("receive from server:%v\n", string(buf[:n]))
	defer conn.Close()
}

func contextHandler(contextR []byte) {
	log.Print("Received context: " + string(contextR))
	spres := bytes.Split(contextR, []byte("%"))
	if len(spres) != 2 {
		log.Print("Proto mismatch!")
		return
	}
	if bytes.Equal(spres[0], []byte("HEARTB")) {
		username := string(spres[1])
		for _, v := range user {
			if username == v {
				log.Print("Old user: " + username)
				pipe <- username
				return
			}
		}
		name := string(spres[1])
		user = append(user, name)
		connected(name)
		go clientHandler(name)
	}
}

func clientHandler(name string) {
	status := 1 //1: connected, 0: disconnect
	for {
		inchan := make(chan int)
		go func() {
			select {
			case <-time.After(12 * time.Second):
				inchan <- 1
			case <-inchan:
				close(inchan)
			}
		}()

	SELECT:
		select {
		case result := <-inchan:
			if result == 1 {
				log.Print(name + " Thread Timeout!")
				close(inchan)
				if status != 0 {
					disconnected(name)
				}
				status = 0
			}
		case nameR := <-pipe:
			if nameR != name {
				pipe <- nameR
				goto SELECT
			}
			if status == 0 {
				reconnected(name)
				status = 1
			}
			inchan <- 1
		}

	}
}

func connected(context0 string) {

}

func reconnected(context0 string) {

}

func disconnected(context0 string) {

}
