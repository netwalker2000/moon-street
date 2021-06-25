package main

import (
	"encoding/binary"
	"encoding/json"
	"moon-street/common"
	"net"
)

func main() {
	var req = common.RpcData{
		Name: "register",
		Args: []interface{}{"bbb", "password2", "zzzzzzz@a.com"},
	}
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
	_, err = conn.Write(buf)
	if err != nil {
		panic(err)
	}
}
