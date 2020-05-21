package apple_conf

import (
	"log"
)

type server struct {
	Host   string
	Port   int
	Labels []string
}

type redis struct {
	Addr string
	Db   int
}

type Config struct {
	Server server
	Redis  redis
}

var conf *Config

func GetAppConfig() *Config {
	if conf == nil {
		log.Panic("not init yet.")
	}
	return conf
}

func InitConfig(c *Config) {
	conf = c
}
