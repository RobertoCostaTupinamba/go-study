package main

import "github.com/RobertoCostaTupinamba/go-study/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
