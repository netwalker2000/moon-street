package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"moon-street/common"
	"net"
)

func main() {
	var req = common.RpcData{
		Name: "register",
		Args: []interface{}{"z15", "password", "xxx@yyyyyy.com"},
	}
	// var req = common.RpcData{
	// 	Name: "login",
	// 	Args: []interface{}{"goon_test_3", "password"},
	// }
	rpcCall(req)
}

func rpcCall(data common.RpcData) {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}
	req, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 4+len(req))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(req)))
	copy(buf[4:], req)
	log.Printf("Send: %s", string(req))
	_, err = conn.Write(buf)
	if err != nil {
		panic(err)
	}
}
