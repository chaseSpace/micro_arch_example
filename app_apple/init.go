package main

import (
	"github.com/leigg-go/go-util/_config"
	"go_project_template/config/apple_conf"
)

func Init() {
	c := new(apple_conf.Config)

	loader := _config.NewLoader("toml")
	loader.SetFileName("apple")
	loader.MustLoad(c)
	apple_conf.InitConfig(c)
}
