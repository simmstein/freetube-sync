package server

import (
	"flag"
	"os"
)

type Config struct {
	BindAddress string
	DbPath      string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}

func (c *Config) Define(bindAddress, dbPath string) {
	c.BindAddress = bindAddress
	c.DbPath = dbPath
}

func InitConfig() {
	dbPath := flag.String("d", os.Getenv("HOME")+"/freetube.sqlite", "Path to SQlite database")
	bindAddress := flag.String("b", ":1323", "Bind address")
	flag.Parse()

	GetConfig().Define(*bindAddress, *dbPath)
}
