package main

import (
	"log"
)

func main() {
	/*
		cmd可以是app的运行时工具，也可以是app的静态辅助工具
	*/
	log.Printf("cmd run success!")
}

/*
Use cmd:
	cd /path/to/go_project_template/cmd/applectl
	go install
	applectl xxx
*/
