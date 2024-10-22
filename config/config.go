package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *viper.Viper

func LoadConfig() {
	AppConfig = viper.New()
	AppConfig.SetConfigName("config")
	AppConfig.SetConfigType("yaml")
	AppConfig.AddConfigPath("./config")

	if err := AppConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig.AutomaticEnv()

	log.Println("Configuration loaded successfully")
}
