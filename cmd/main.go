package main

import (
	"bufio"
	"log"
	"moon-street/common"
	"moon-street/internal/controller"
	"net"
)

func main() {
	log.Println("Begin...")
	PORT := ":8001"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	log.Println("Ready")

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for {
		cmd, ok := common.GetStringLine(scanner)
		controller.Route(cmd)
		if !ok {
			return
		}

	}
}
