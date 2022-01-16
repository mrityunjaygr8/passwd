package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST string
	DB_PORT int
	DB_NAME string
	DB_USER string
	DB_PASS string
	PORT    int
	HOST    string
}

func GetConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.toml")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "passwd")
	viper.SetDefault("PORT", 1337)
	viper.SetDefault("HOST", "localhost")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config: %w\n", err))
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return c
}
