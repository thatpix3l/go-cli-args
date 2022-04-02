package main

import (
	"github.com/thatpix3l/args/app"
	"github.com/thatpix3l/args/cmd"
)

func main() {

	config := cmd.GenerateConfig()
	app.Start(&config)

}
