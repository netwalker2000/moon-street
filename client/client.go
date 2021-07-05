package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"math/rand"
	"moon-street/common"
	"net"
	"strconv"
	"time"
)

func main() {
	//random uid
	timestamp := time.Now().UnixNano()
	log.Println(timestamp)
	r := rand.New(rand.NewSource(timestamp))
	randUname := "u" + strconv.Itoa(r.Intn(9999999)+50)
	log.Println(randUname)

	var req = common.RpcData{
		Name: "login",
		Args: []interface{}{"u51", "password"},
	}
	// var req = common.RpcData{
	// 	Name: "login",
	// 	Args: []interface{}{"goon_test_3", "password"},
	// }
	rpcCall([]common.RpcData{req, req, req, req, req})
}

func rpcCall(data []common.RpcData) {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		panic(err)
	}
	req, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	log.Printf("body: %v", req)
	buf := make([]byte, 4+len(req))
	len_header := len(req)
	binary.BigEndian.PutUint32(buf[:4], uint32(len_header))
	log.Printf("header len: %d", len_header)
	copy(buf[4:], req)
	log.Printf("Send: %s", string(req))
	_, err = conn.Write(buf)
	if err != nil {
		panic(err)
	}
}
