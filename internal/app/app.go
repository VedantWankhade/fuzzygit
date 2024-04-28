package app

import (
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/add"
	branch "github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/checkout"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/diff"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/help"
)

const (
	HELP     Command = "help"
	CHECKOUT Command = "checkout"
	DIFF     Command = "diff"
	ADD      Command = "add"
)

func Run(command Command, flags []string) {
	switch command {
	case HELP:
		help.Invoke()
	case CHECKOUT:
		branch.Invoke(flags)
	case DIFF:
		diff.Invoke(flags)
	case ADD:
		add.Invoke(flags)
	}
}
