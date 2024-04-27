package main

import (
	"os"

	"github.com/vedantwankhade/fuzzygit/internal/app"
)

func main() {
	var cmd app.Command
	app.InfoLogger.Println("fuzzygit invoked")
	if len(os.Args) < 2 {
		cmd = app.HELP
	} else {
		switch os.Args[1] {
		case "help":
			cmd = app.HELP
		case "checkout":
			cmd = app.CHECKOUT
		default:
			cmd = app.HELP
		}
	}
	app.InfoLogger.Println("running command", cmd)
	app.Run(cmd)
}
