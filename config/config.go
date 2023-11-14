package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type ENV string

const (
	ENV_RELEASE ENV = `release`
	ENV_TEST    ENV = `test`
	ENV_DEV     ENV = `dev`
)

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		var err error
		var configPath = "config.dev.yaml"
		var env = os.Getenv("ENV")
		if env == string(ENV_RELEASE) {
			configPath = "config.release.yaml"
		}
		if config, err = LoadConfig(configPath); err != nil {
			panic(fmt.Sprintf("load config failed for %v", err))
		}
	})
	return config
}

type Config struct {
	SVNPath string `mapstructure:"svn" desc:"svn 客户端地址"`
}

func LoadConfig(configPath string) (*Config, error) {
	// Set up viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// Read in the config file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal the config into a struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
