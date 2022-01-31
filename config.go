package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	TGKey      string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func CreateConfig() *Config {
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.base.yml")
	if err != nil {
		panic(err)
	}
	err = config.LoadFiles("config.custom.yml")

	if err != nil {
		panic(err)
	}
	c := Config{}
	err = config.BindStruct("", &c)

	return &c
}
