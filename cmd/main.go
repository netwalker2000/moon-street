package main

import (
	"bufio"
	"fmt"
	"log"
	"moon-street/internal/controller"
	"net"
)

func main() {
	fmt.Println("Begin...")
	PORT := ":8001"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Println("Ready")

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
		cmd, ok := getStringLine(scanner)
		controller.Route(cmd)
		if !ok {
			return
		}

	}
}
func getStringLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		fmt.Println(scanner.Err())
		return "", false
	}
	return scanner.Text(), true
}
