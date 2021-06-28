package util

import (
	"crypto/sha256"
	"fmt"
	"moon-street/config"
)

func encrypt(password string, salt string) string {
	saltedPassword := password + salt
	ret := fmt.Sprintf("%x", sha256.Sum256([]byte(saltedPassword)))
	return ret
}

func EncryptWithSalt(password string) string {
	salt := config.ConfigSingleton.Server.Salt
	return encrypt(password, salt)
}
