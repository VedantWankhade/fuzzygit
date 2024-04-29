package checkout

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
)

func Invoke(flags []string) {
	if slices.Contains(flags, "-c") {
		type commit struct {
			commit string
		}
		logMap := make(map[string]commit)
		var hashes []string
		// flags = append(flags, "--oneline")
		res := git.Cmd("log", "--oneline")
		logLines := strings.Split(res, "\n")
		for _, l := range logLines {
			arr := strings.SplitAfterN(l, " ", 2)
			hash := strings.TrimSpace(arr[0])
			hashes = append(hashes, hash)
			logMap[hash] = commit{
				commit: arr[1],
			}
		}
		f, err := fzf.New()
		if err != nil {
			log.Fatal(err)
		}

		idxs, err := f.Find(
			logLines,
			func(i int) string { return hashes[i] + " " + logMap[hashes[i]].commit },
			fzf.WithPreviewWindow(func(i, width, height int) string {
				filesChanged := git.Cmd("diff", hashes[i], "--name-only")
				return fmt.Sprintf("Commit message: %s\n\nFiles changed:\n%v", logMap[hashes[i]].commit, filesChanged)
			}),
		)
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range idxs {
			fmt.Println(logMap[hashes[i]].commit)
			fmt.Println(hashes[i])
		}
		git.Cmd("checkout", hashes[idxs[0]])
		git.Cmd("switch", "-c", ("branch-" + hashes[idxs[0]]))
	} else if slices.Contains(flags, "-t") {
		t := git.Cmd("tag", "-l")
		tags := strings.Split(string(t), "\n")
		f, err := fzf.New()
		if err != nil {
			log.Fatal(err)
		}
		idxs, err := f.Find(tags, func(i int) string { return tags[i] })
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range idxs {
			git.Cmd("checkout", tags[i])
			git.Cmd("switch", "-c", ("tag-" + tags[i]))
		}
	} else {
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
		}
	}
}
