package common

import (
	"bufio"
	"log"
)

type RpcData struct {
	Name string        //remote function name
	Args []interface{} //remote function args
}

func GetStringLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		log.Println(scanner.Err())
		return "", false
	}
	return scanner.Text(), true
}
