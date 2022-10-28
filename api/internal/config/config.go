package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config store all configuration of the application
// The value are readby viper from the congfig file
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

const (
	// EnvPath Path env file
	EnvPath = "."
	// ConfigName env file name
	ConfigName = "app"
	// ConfigExtension  env file extension
	ConfigExtension = "env"
)

// LoadConfig config value from env file
func LoadConfig() (Config, error) {
	viper.AddConfigPath(EnvPath)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigExtension)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	// re-validate
	if config.DBDriver == "" {
		log.Fatal("Cannot load DBDriver from config file.")
	}

	if config.ServerAddress == "" {
		log.Fatal("Cannot load ServerAddress from config file.")
	}

	if config.DBSource == "" {
		log.Fatal("Cannot load DB_URL from config file.")
	}

	return config, nil
}
