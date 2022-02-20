package main

import (
	"live-chat/internal/app"
)

func main() {
	serv := app.NewServer()
	serv.Start()
}
