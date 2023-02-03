package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Server struct {
	Port string `mapstructure:"PORT"`
}

type DB struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"` 
	Name string `mapstructure:"NAME"`
	Driver string `mapstructure:"DRIVER"`
}

type RabbitMQ struct {
	Protocol string `mapstructure:"PROTOCOL"`
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"` 
	VHost string `mapstructure:"VHOST"` 
}

type Config struct {
	Server Server `mapstructure:"SERVER"`
	DB DB `mapstructure:"DB"`
	RabbitMQ RabbitMQ `mapstructure:"RABBITMQ"`
}

func LoadEnv(path string) (*Config) {
	replacer := strings.NewReplacer(".", "_")
    	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)               // optionally look for config in the working directory
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading env, %s\n", err)
	}

	var env Config 
	err = viper.Unmarshal(&env)
	return &env
}
