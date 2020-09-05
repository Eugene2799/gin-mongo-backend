package main

import (
	"gin-mongo-backend/app"
	"flag"
)

func main() {
	configFile := flag.String("c", "config.yaml", "Config file")
	flag.Parse()
	app.Run(*configFile)
}
