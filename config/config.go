package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Listen           string
	WriteTimeout     time.Duration
	IdleTimeout      time.Duration
	ReadTimeout      time.Duration
	GracefulShutdown time.Duration
}

type DataService struct {
	APIKey string
	URL    string
}
type Configuration struct {
	Database    Database
	DataService DataService
	Server      Server
}

type Database struct {
	DSN            string
	MigrationsPath string
}

func InitConfig(configFile string) (*Configuration, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	v.SetEnvPrefix("auth")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Configuration{}

	err = v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
