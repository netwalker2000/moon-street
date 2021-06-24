package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	InitConfig()
	summary := viper.GetViper().AllKeys()
	fmt.Println(summary)
	port := viper.GetString("database.port")
	fmt.Println(port)
}
