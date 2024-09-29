package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

//var EnvCfg envConfig

var EnvCfg = struct {
	ServerPort     string `env:"SERVER_PORT" envDefault:"80"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"debug"`
	AutoMigrate    bool   `env:"AUTO_MIGRATE" envDefault:"true"`
	LocationKey    string `env:"LOCATION_KEY"`
	LocationSecret string `env:"LOCATION_SECRET"`

	// 七牛云配置
	QiNiuAccessKey string `env:"QI_NIU_ACCESS_KEY"`
	QiNiuSecretKey string `env:"QI_NIU_SECRET_KEY"`
}{}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Can not read env from file system, please check the right this program owned.")
	}

	//EnvCfg = envConfig{}

	if err := env.Parse(&EnvCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}

	println(EnvCfg.ServerPort)
}
