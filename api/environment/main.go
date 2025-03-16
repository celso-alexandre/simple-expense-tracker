package environment

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type Config struct {
	DatabaseURL string `envconfig:"DIRECT_URL" required:"true"`
}

func GetConfig() Config {
	err := gotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
