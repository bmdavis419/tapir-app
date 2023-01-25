package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	MONGODB_URI  string `mapstructure:"MONGODB_URI"`
	MONGODB_NAME string `mapstructure:"MONGODB_NAME"`
	PORT         string `mapstructure:"PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			MONGODB_URI:  os.Getenv("MONGODB_URI"),
			MONGODB_NAME: os.Getenv("MONGODB_NAME"),
			PORT:         os.Getenv("PORT"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// validate config here
	if config.MONGODB_URI == "" {
		err = errors.New("MONGODB_URI is required")
		return
	}

	if config.MONGODB_NAME == "" {
		err = errors.New("MONGODB_NAME is required")
		return
	}

	return
}
