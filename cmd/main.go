package main

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
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
	if !config.ConfigSingleton.Debug {
		log.SetOutput(ioutil.Discard)
	}
	addr := config.ConfigSingleton.Server.Address
	log.Printf("Prepare to listen on %s", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("cannot net.Listen, please check port!")
		panic(err)
	}
	defer listener.Close()
	log.Println("Ready")

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Printf("cannot listener.Accept, please check!")
			log.Fatal(err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	var (
		err_resp = []byte{0, 0, 0, 0}
		pos_resp = []byte{0, 0, 0, 10, 123, 34, 99, 111, 100, 101, 34, 58, 48, 125}
	)
	header := make([]byte, 4)
	if _, err := c.Read(header); err != nil {
		log.Printf("error when read conn header, end this conn :%v", err)
		c.Write(err_resp)
		return
	}
	bodyLen := binary.BigEndian.Uint32(header)
	//log.Printf("len of body(big endian): %d", bodyLen)
	body := make([]byte, int(bodyLen))
	if _, err := c.Read(body); err != nil {
		log.Printf("error when read conn body, end this conn :%v", err)
		c.Write(err_resp)
		return
	}
	var rData []common.RpcData
	if err := json.Unmarshal(body, &rData); err != nil {
		log.Printf("error when convert remote data, end this conn :%v", err)
		c.Write(err_resp)
		return
	}
	defer func() {
		if e := recover(); e != nil {
			log.Printf("error when goroutine: %v", e)
		}
	}()
	for i, v := range rData {
		err := controller.Route(v)
		//log.Printf("%v", v)
		if err != nil {
			log.Printf("error when exec No.%d in the batch :%v", i, err)
			c.Write(err_resp)
			return
		}
	}
	c.Write(pos_resp)
}
