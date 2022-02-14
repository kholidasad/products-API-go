package main

import (
	"github.com/kholidasad/products-API-go/config"
	"github.com/kholidasad/products-API-go/app"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")	
}