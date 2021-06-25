package common

import (
	"bufio"
	"log"
)

func GetStringLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		log.Println(scanner.Err())
		return "", false
	}
	return scanner.Text(), true
}
