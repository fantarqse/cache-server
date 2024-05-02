package config

import "github.com/spf13/viper"

type Config struct {
	Server Server
	Redis  Redis
}

type Server struct {
	Host string
	Port string
}

type Redis struct {
	Host string
	Port string
}

func Read() (Config, error) {
	var cfg Config

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config/")
	viper.SetConfigName("local")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
