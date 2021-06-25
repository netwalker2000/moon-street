package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	configFile = "/Users/zezhou.zheng/gospace/moon-street/config/config.yml"
	configType = "yml"
)

var ConfSingleton Config = newConfig()

type (
	Config struct {
		Debug    bool
		Server   Server
		Database Database
	}

	Server struct {
		Address string
		Salt    string
	}

	Database struct {
		Driver   string
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}
)

func newConfig() Config {
	initConfig()
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fatalErr := fmt.Errorf("error when new config! origin: %v", err)
		log.Fatalf("fatal error: %v", fatalErr)
		os.Exit(1)
	}
	log.Printf("Successfully init config: %v", conf)
	return *conf
}

func initConfig() {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig() //todo: password can not in git file

	if err != nil {
		fatalErr := fmt.Errorf("error when init config! origin: %v", err)
		log.Fatalf("fatal error: %v", fatalErr)
		os.Exit(1)
	}
}
