package checkout

import (
	"log"
	"os/exec"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
)

func Invoke(flags []string) {
	b := git.Cmd("branch", flags...)
	arr := strings.Split(string(b), "\n")
	var branches []string
	for _, i := range arr {
		b := strings.TrimSpace(i)
		branches = append(branches, strings.Replace(b, "* ", "", 1))
	}
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}
	idxs, err := f.Find(branches, func(i int) string { return branches[i] })
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range idxs {
		cmd := exec.Command("git", "checkout", branches[i])
		cmd.Run()
	}
}
