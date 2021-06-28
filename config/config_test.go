package config

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	summary := viper.GetViper().AllKeys()
	log.Println(summary)
	port := viper.GetString("database.port")
	log.Println(port)
}
