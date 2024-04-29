package rename

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/input"
)

func Invoke(flags []string) {
	subcmd := "checkout"
	b := git.Cmd("branch", flags...)
	arr := strings.Split(string(b), "\n")
	var branches []string
	for _, i := range arr {
		if strings.Contains(i, "->") {
			continue
		}
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
	if slices.Contains(flags, "-r") {
		subcmd = "switch"
		for _, i := range idxs {
			branches[i] = strings.SplitAfterN(branches[i], "/", 2)[1]
		}
		git.Cmd("fetch")
	}
	for _, i := range idxs {
		git.Cmd(subcmd, branches[i])
		ch := make(chan string)
		go input.Get(ch)
		name := <-ch
		fmt.Println(branches[i], "->", name)
		// git branch -m <oldname> <newname>
		git.Cmd("branch", "-m", branches[i], name)
	}
}
