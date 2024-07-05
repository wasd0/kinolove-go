package config

import (
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
)

const (
	EnvDev   = "dev"
	EnvProd  = "prod"
	EnvStage = "stage"
)

const envConfigPath = "CONFIG_PATH"

type Config struct {
	Env    string  `yaml:"env" env-default:"prod" env-required:"true"`
	DbPath string  `yaml:"db_path" env-required:"true"`
	Server *Server `yaml:"server" env-required:"true"`
}

type Server struct {
	Port        string        `yaml:"port" env-default:"8100"`
	Host        string        `yaml:"host" env-default:"localhost"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"69s"`
}

var isLoaded = false

func MustRead() *Config {

	if !isLoaded {
		isLoaded = true
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	configPath := os.Getenv(envConfigPath)

	if configPath == "" {
		log.Fatalf("%s is empty", envConfigPath)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists by path: %s", configPath)
	}

	config := Config{}

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal("Error while config read")
	}

	return &config
}
