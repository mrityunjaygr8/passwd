package utils

import (
	"fmt"
	"log"
	"regexp"
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
	DB_SSL  bool
}

func GetConfig() Config {
	viper.BindEnv("db_host")
	viper.BindEnv("db_port")
	viper.BindEnv("db_name")
	viper.BindEnv("db_user")
	viper.BindEnv("db_pass")
	viper.BindEnv("db_ssl")
	viper.BindEnv("port")
	viper.BindEnv("host")
	viper.BindEnv("database_url")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.toml")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "passwd")
	viper.SetDefault("PORT", 1337)
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("DB_SSL", false)

	err := viper.ReadInConfig()
	if err != nil {
		if ok := strings.Contains(err.Error(), "no such file or directory"); ok {
			u := viper.Get("DB_USER")
			p := viper.Get("DB_PASS")

			// This is for heroku only
			d := viper.Get("DATABASE_URL")
			// fmt.Println(fmt.Sprint(d))
			if (u == nil || p == nil) && d == nil {
				panic(fmt.Errorf("Fatal error config: %w\n", fmt.Errorf("You have not specified a config file or used envs for DB_USER and DB_PASS")))
			}

			// This is for heroku only
			if d != nil {
				// r, _ := regexp.Compile(`postgres://([\w]+):([\w]+)@([\w]+):([\w]+)/([\w]+)`)
				r, _ := regexp.Compile(`postgres://(?P<user>[\w]+):(?P<pass>[\w]+)@(?P<host>[\w-.]+):(?P<port>[\d]+)/(?P<name>[\w]+)`)
				matches := r.FindSubmatch([]byte(fmt.Sprint(d)))
				viper.Set("DB_USER", string(matches[1]))
				viper.Set("DB_PASS", string(matches[2]))
				viper.Set("DB_HOST", string(matches[3]))
				viper.Set("DB_PORT", string(matches[4]))
				viper.Set("DB_NAME", string(matches[5]))
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

	fmt.Println(c)

	return c
}
