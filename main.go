package main

import (
	"github.com/thatpix3l/args/app"
	"github.com/thatpix3l/args/cmd"
	"github.com/thatpix3l/args/config"
)

func main() {

	var config config.Config
	cmd.GenerateConfig(&config)
	app.Start(&config)

}
