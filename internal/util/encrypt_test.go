package util

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	encrypted := encrypt("password", "CAFEBABE")
	if encrypted == "cb861abbe31ab241bc12006e24227eb6ce5e9b28c44e3c4463845825a6a4cd48" {
		log.Println("Success")
	} else {
		t.FailNow()
	}
}
