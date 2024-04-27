package branch

import (
	"log"
	"os/exec"
	"strings"

	"github.com/koki-develop/go-fzf"
)

func GetBranches() []string {
	cmd := exec.Command("git", "branch", "--list")
	b, _ := cmd.Output()
	b = []byte(strings.TrimSpace(string(b)))
	arr := strings.Split(string(b), "\n")
	var branches []string
	for _, i := range arr {
		b := strings.TrimSpace(i)
		branches = append(branches, strings.Replace(b, "* ", "", 1))
	}
	return branches
}

func Invoke() {
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
