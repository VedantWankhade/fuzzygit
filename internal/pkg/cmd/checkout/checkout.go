package checkout

import (
	"log"
	"os/exec"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
)

func GetBranches() []string {
	b := git.Cmd("branch")
	arr := strings.Split(string(b), "\n")
	var branches []string
	for _, i := range arr {
		b := strings.TrimSpace(i)
		branches = append(branches, strings.Replace(b, "* ", "", 1))
	}
	return branches
}

func Invoke(flags []string) {
	items := GetBranches()
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}
	idxs, err := f.Find(items, func(i int) string { return items[i] })
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range idxs {
		cmd := exec.Command("git", "checkout", items[i])
		cmd.Run()
	}
}
