package storage

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type DataBase struct {
	USER     string
	PASSWORD string
	NAME     string
	HOST     string
	PORT     string
}

type Config struct {
	Database DataBase
}

func GetConfig() Config {
	filename := flag.String("config", "./config.toml", "Name of the config file")
	_, err := os.Stat(*filename)
	if err != nil {
		log.Fatal("Config file is missing: ", filename)
	}

	var config Config
	if _, err := toml.DecodeFile(*filename, &config); err != nil {
		log.Fatal(err)
	}
	log.Print(config)
	return config
}
