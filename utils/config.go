package utils

import (
	"fmt"
	"log"
	"strings"

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
	viper.SetEnvPrefix("passwd")
	viper.BindEnv("db_host")
	viper.BindEnv("db_port")
	viper.BindEnv("db_name")
	viper.BindEnv("db_user")
	viper.BindEnv("db_pass")
	viper.BindEnv("port")
	viper.BindEnv("host")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.toml")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "passwd")
	viper.SetDefault("PORT", 1337)
	viper.SetDefault("HOST", "localhost")

	err := viper.ReadInConfig()
	if err != nil {
		if ok := strings.Contains(err.Error(), "no such file or directory"); ok {
			u := viper.Get("DB_USER")
			p := viper.Get("DB_PASS")
			if u == nil || p == nil {
				panic(fmt.Errorf("Fatal error config: %w\n", fmt.Errorf("You have not specified a config file or used envs for DB_USER and DB_PASS")))
			}
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config: %w\n", err))
		}
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return c
}
