package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Client Begin...")
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("Failed")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		conn.Write([]byte(input))
		//server
		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("Failed %s\n", err)
			continue
		}

		fmt.Print("reply:" + string(buf[0:cnt]))
	}

}
