package app

import (
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/branch"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/help"
)

const (
	HELP   Command = "help"
	BRANCH Command = "branch"
	ADD    Command = "add"
)

func Run(command Command) {
	switch command {
	case HELP:
		help.Invoke()
	case BRANCH:
		branch.Invoke()
	}
}
