package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfiguration() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/homeservices")
	viper.AddConfigPath("$HOME/.homeservices")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	return nil
}
