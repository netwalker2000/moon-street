package util

import (
	"crypto/sha256"
	"fmt"
	"moon-street/config"
)

func EncryptWithSalt(password string) string {
	salt := config.ConfigSingleton.Server.Salt
	saltedPassword := password + salt
	ret := fmt.Sprintf("%x", sha256.Sum256([]byte(saltedPassword)))
	return ret
}
