package client

import (
	"flag"
	"os"
	"strings"
)

type Config struct {
	Server   string
	Path     string
	Hostname string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}

func (c *Config) Define(server, hostname, path string) {
	c.Server = strings.TrimRight(server, "/")
	c.Hostname = hostname
	c.Path = path
}

func InitConfig() {
	defaultHostname, _ := os.Hostname()

	path := flag.String("p", os.Getenv("HOME")+"/.config/FreeTube", "Path to FreeTube config directory")
	hostname := flag.String("h", defaultHostname, "Hostname")
	server := flag.String("s", "", "Server to sync")
	flag.Parse()

	GetConfig().Define(*server, *hostname, *path)
}
