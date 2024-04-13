package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort  string `mapstructure:"SERV_PORT"`
	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresUser string `mapstructure:"POSTGRES_USER"`
	PostgresPass string `mapstructure:"POSTGRES_PASS"`
	PostgresName string `mapstructure:"POSTGRES_NAME"`
	PostgresPort string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
