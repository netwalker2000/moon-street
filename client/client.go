package main

import (
	"bufio"
	"log"
	"moon-steet/common"
	"moon-street/common"
	"net"
	"strings"
)

func main() {
	log.Println("Client Begin...")
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Println("Failed")
		return
	}
	reader := bufio.NewScanner(conn)

	for {
		log.Println(reader)
		input := common.GetStringLine()
		input = strings.TrimSpace(input)
		log.Printf("input: %s", input)
		conn.Write([]byte(input))

	}

}
