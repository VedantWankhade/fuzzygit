package app

import (
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/add"
	branch "github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/checkout"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/diff"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/help"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/log"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/rename"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/unstage"
)

const (
	HELP     Command = "help"
	CHECKOUT Command = "checkout"
	DIFF     Command = "diff"
	ADD      Command = "add"
	UNSTAGE  Command = "unstage"
	LOG      Command = "log"
	RENAME   Command = "rename"
)

func Run(command Command, flags []string) {
	switch command {
	case HELP:
		help.Invoke(flags)
	case CHECKOUT:
		branch.Invoke(flags)
	case DIFF:
		diff.Invoke(flags)
	case ADD:
		add.Invoke(flags)
	case UNSTAGE:
		unstage.Invoke()
	case LOG:
		log.Invoke(flags)
	case RENAME:
		rename.Invoke(flags)
	}
}
