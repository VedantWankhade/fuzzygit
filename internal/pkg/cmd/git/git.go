package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vedantwankhade/fuzzygit/internal/config"
)

func Cmd(subcommand string, flags ...string) string {
	config.InfoLogger.Println("running git", subcommand, strings.Join(flags, " "))
	var options []string
	options = append(options, subcommand)
	options = append(options, flags...)
	cmd := exec.Command("git", options...)
	res, err := cmd.Output()
	if err != nil {
		config.ErrorLogger.Println("error running git", strings.Join(options, " "), err)
		fmt.Println("error running git", strings.Join(options, " "), err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(res))
}
