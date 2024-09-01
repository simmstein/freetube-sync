package server

import (
	"flag"
	"os"

	"gitnet.fr/deblan/freetube-sync/logger"
)

type Config struct {
	BindAddress string
	DbPath      string
	LogLevel    int
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}

func (c *Config) Define(bindAddress, dbPath string, logErrorLevel, logWarnLevel, logInfoLevel bool) {
	c.BindAddress = bindAddress
	c.DbPath = dbPath

	if logInfoLevel {
		c.LogLevel = int(logger.Info)
	} else if logErrorLevel {
		c.LogLevel = int(logger.Error)
	} else if logWarnLevel {
		c.LogLevel = int(logger.Warn)
	} else if logInfoLevel {
		c.LogLevel = int(logger.Info)
	} else {
		c.LogLevel = int(logger.Silent)
	}
}

func InitConfig() {
	dbPath := flag.String("d", os.Getenv("HOME")+"/freetube.sqlite", "Path to SQlite database")
	bindAddress := flag.String("b", ":1323", "Bind address")
	logErrorLevel := flag.Bool("v", false, "Log errors")
	logWarnLevel := flag.Bool("vv", false, "Log warns")
	logInfoLevel := flag.Bool("vvv", false, "Log infos")
	flag.Parse()

	GetConfig().Define(*bindAddress, *dbPath, *logErrorLevel, *logWarnLevel, *logInfoLevel)
}
