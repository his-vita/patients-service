package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
	Sql    Sql    `yaml:"sql"`
}

type Server struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-default:"8080"`
}

type Db struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	DbName   string `yaml:"db_name" env-default:"patients_db"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

type Sql struct {
	Path string `yaml:"path" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
