package config

import (
	"log"

	"github.com/spf13/viper"
)

type DB struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"` 
	Name string `mapstructure:"NAME"`
	Driver string `mapstructure:"DRIVER"`
}

type RabbitMQ struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"` 
}

type Config struct {
	DB DB `mapstructure:"DB"`
	RabbitMQ RabbitMQ `mapstructure:"RABBIT_MQ"`
}

func LoadEnv(path string) (*Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading env, %s\n", err)
	}

	var env Config 
	err = viper.Unmarshal(&env)
	return &env
}
