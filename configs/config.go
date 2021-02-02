package config

import (
	"log"

	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Gin *GinConfig
}

type GinConfig struct {
	Host string
	Port string
	Mode string
}

func NewConfig() *Config {
	configPath := "./"
	runPath, _ := os.Getwd()
	matchPathStatus := false
	pathArr := strings.Split(runPath, "/")
	for i := len(pathArr) - 1; i > 0; i-- {
		configPath += "../"
		if pathArr[i] == "cmd" || pathArr[i] == "test" || pathArr[i] == "migration" {
			matchPathStatus = true
			break
		}
	}
	if !matchPathStatus {
		configPath = "./"
	}
	configPath += "configs"

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Gin: &GinConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
			Mode: viper.GetString("server.mode"),
		},
	}
}
