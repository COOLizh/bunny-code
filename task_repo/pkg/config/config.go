// Package config returns configuration parameters
package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	DatabaseURL     string `mapstructure:"DATABASE_URL"`
	APIPort         string `mapstructure:"API_PORT"`
	FileLogName     string `mapstructure:"FILE_LOG"`
	JwtSalt         string `mapstructure:"JWT_SALT"`
	ExecutionerPort string `mapstructure:"EXECUTIONER_PORT"`
	ClientAddr      string `mapstructure:"CLIENT_ADDR"`
}

// LoadConfig is method for populate Config struct
func (config *Config) LoadConfig() error {
	viper.AddConfigPath("../../")
	viper.AddConfigPath(".")

	viper.SetConfigName("default")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err)
	}

	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.MergeInConfig(); err != nil {
		log.Print(err)
		_ = viper.BindEnv("DATABASE_URL")
		_ = viper.BindEnv("API_PORT")
		_ = viper.BindEnv("FILE_LOG")
		_ = viper.BindEnv("JWT_SALT")
		_ = viper.BindEnv("EXECUTIONER_PORT")
		_ = viper.BindEnv("CLIENT_ADDR")
	}

	err = viper.Unmarshal(&config)
	return err
}
