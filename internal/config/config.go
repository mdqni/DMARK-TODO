package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `env:"ENV" env-default:"local"`
	ConnString string `env:"CONN_STRING" env-default:"postgres://postgres:1234@localhost:5433/tododb?sslmode=disable"`

	App AppConfig
}

type AppConfig struct {
	Title  string `env:"APP_TITLE" env-default:"To-Do App"`
	Width  int    `env:"APP_WIDTH" env-default:"1024"`
	Height int    `env:"APP_HEIGHT" env-default:"768"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("cannot read config from env: %s", err)
	}
	return &cfg
}
