package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Sql struct {
	Host string
	Port int
	DB   string
}

type Redis struct {
	Host string
	Port int
}

type Config struct {
	Sql   Sql
	Redis Redis
}

var config Config

func InitConfig(configName string) error {
	viper.AddConfigPath("./shared/config/secret/")
	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return errors.New("Error loading .json file (not found)")
		} else {
			// Config file was found but another error was produced
			return errors.New("Error loading .json file (found but error)")
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return errors.New("Error unmarshal .json file")
	}

	return nil
}

func GetConfig() Config {
	return config
}
