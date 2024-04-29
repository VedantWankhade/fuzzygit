package main

import (
	"os"

	"github.com/vedantwankhade/fuzzygit/internal/app"
	"github.com/vedantwankhade/fuzzygit/internal/config"
)

func main() {
	var cmd app.Command
	var flags []string
	config.InfoLogger.Println("fuzzygit invoked")
	if len(os.Args) < 2 {
		cmd = app.HELP
	} else {
		switch os.Args[1] {
		case "help":
			cmd = app.HELP
		case "checkout":
			cmd = app.CHECKOUT
		case "diff":
			cmd = app.DIFF
		case "add":
			cmd = app.ADD
		case "unstage":
			cmd = app.UNSTAGE
		case "log":
			cmd = app.LOG
		case "rename":
			cmd = app.RENAME
		default:
			cmd = app.HELP
		}
	}
	config.InfoLogger.Println("running command", cmd)
	if cmd == app.HELP {
		flags = []string{}
	} else {
		flags = os.Args[2:]
	}
	app.Run(cmd, flags)
}
