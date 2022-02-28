package main

import (
	"flag"
	"live-chat/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := app.NewConfig()

	serv := app.NewServer(config)
	serv.Start()
}
