package config

import "github.com/jinzhu/configor"

type ServerConfig struct {
	Port int `yaml:"port"`
	ListenAddr string `yaml:"listenAddr"`
}


type Config struct {
	Server ServerConfig `yaml:"server"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := configor.Load(cfg, "config.yml"); err != nil {
		return nil, err
	}
	return cfg, nil
}