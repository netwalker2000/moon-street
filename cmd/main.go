package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"moon-street/common"
	"moon-street/config"
	"moon-street/internal/controller"
	"moon-street/internal/di"
	"net"
)

func main() {
	log.Println("Begin...")
	di.InitDependenciesUseFactories()
	addr := config.ConfigSingleton.Server.Address
	log.Printf("Prepare to listen on %s", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("Ready")

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	header := make([]byte, 4)
	if _, err := c.Read(header); err != nil {
		log.Printf("error when read conn header, end this conn :%v", err)
		return
	}
	bodyLen := binary.BigEndian.Uint32(header)
	body := make([]byte, int(bodyLen))
	if _, err := c.Read(body); err != nil {
		log.Printf("error when read conn body, end this conn :%v", err)
		return
	}
	var rData common.RpcData
	if err := json.Unmarshal(body, &rData); err != nil {
		log.Printf("error when convert remote data, end this conn :%v", err)
		return
	}
	controller.Route(rData)
}
