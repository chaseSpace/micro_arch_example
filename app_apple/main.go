package main

import (
	"go_project_template/app_apple/internal"
	"go_project_template/config/apple_conf"
	"log"
	"time"
)

func main() {
	Init()

	conf := apple_conf.GetAppConfig()
	log.Printf("%+v", *conf)

	log.Println(internal.AppleHi())

	log.Println("apple is running~")

	time.Sleep(2 * time.Second)

	log.Println("apple exit~")
}
