package main

import (
	"flag"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/belyakoff-nikolay/otus-microservices/hw06/internal/app/apiserver"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "configuration file")
	flag.Parse()
}

func main() {
	config := apiserver.NewConfig()
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(b, &config); err != nil {
		panic(err)
	}
	config.DatabaseURL = os.Getenv("DATABASE_URL")

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
