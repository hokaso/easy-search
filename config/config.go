package config

import (
	"easy-search/pkgs/path"
	"fmt"
	"github.com/jinzhu/configor"
)

var (
	DbConfig = struct {
		Database string `required:"true" env:"DBDatabase"`
		Username string `default:"root" env:"DBUsername"`
		Password string `required:"false" env:"DBPassword"`
		Host     string `default:"localhost" env:"DBHost"`
		Port     string `default:"3306" env:"DBPort"`
	}{}

	RedisConfig = struct {
		Db   string `required:"true" default:"0"`
		Host string `required:"true" default:"localhost"`
		Port string `required:"true" default:"6379"`
	}{}
)

func LoadConfig() {
	LoadDbConfig()
	LoadRedisConfig()
}

func LoadDbConfig() {
	err := configor.Load(&DbConfig, fmt.Sprintf("%s/config/database.yml", path.ConfigPath))
	if err != nil {
		panic(err)
	}
}

func LoadRedisConfig() {
	err := configor.Load(&RedisConfig, fmt.Sprintf("%s/config/redis.yml", path.ConfigPath))
	if err != nil {
		panic(err)
	}
}
