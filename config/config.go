package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

const (
	configFile = "/gospace/moon-street/config/config.yml"
	configType = "yml"
)

var ConfigSingleton Config = newConfig()

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
		log.Fatalf("fatal error when read config: %v", err)
		os.Exit(1)
	}
	log.Printf("Successfully init config: %v", conf)
	return *conf
}

func initConfig() {
	homePath := os.Getenv("HOME")
	viper.SetConfigType(configType)
	viper.SetConfigFile(homePath + configFile)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("fatal error when init config: %v", err)
		os.Exit(1)
	}
}
