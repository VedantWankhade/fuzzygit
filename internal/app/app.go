package app

import (
	branch "github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/checkout"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/help"
)

const (
	HELP     Command = "help"
	CHECKOUT Command = "checkout"
	ADD      Command = "add"
)

func Run(command Command) {
	switch command {
	case HELP:
		help.Invoke()
	case CHECKOUT:
		branch.Invoke()
	}
}
