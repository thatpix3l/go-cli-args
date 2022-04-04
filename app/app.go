package app

import (
	"fmt"

	"github.com/thatpix3l/args/config"
)

type color_struct struct {
	blue   string
	green  string
	yellow string
	red    string
	black  string
	reset  string
}

var (
	color = color_struct{
		black:  "\u001b[30m",
		red:    "\u001b[31m",
		green:  "\u001b[32m",
		yellow: "\u001b[33m",
		blue:   "\u001b[34m",
		reset:  "\u001b[0m",
	}
)

// Add color codes to beginning of a given string, reset after end
func colorize(msg_color string, msg string) string {

	return msg_color + msg + color.reset

}

// Beginning of actual program after loading and parsing config files
func Start(c *config.Config) {

	name := "Joe"
	if *c.UseColor {
		name = colorize(color.red, name)
	}
	fmt.Println(name)

	if *c.ShowFunny {
		fmt.Println("LMAO, IMAGINE")
	}

	fmt.Printf("%s\n", c.CoolString)

}
