package main

import (
	"github.com/Technautics/GO-FIG/app"
	"github.com/Technautics/GO-FIG/config"
)

func main() {

	cfg := config.LoadConfig()
	app.StartApp(cfg)

}
