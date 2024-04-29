package log

import (
	"fmt"
	"log"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/pager"
)

func Invoke(flags []string) {
	type commit struct {
		commit string
	}
	logMap := make(map[string]commit)
	var hashes []string
	flags = append(flags, "--oneline")
	res := git.Cmd("log", flags...)
	logLines := strings.Split(res, "\n")
	for _, l := range logLines {
		arr := strings.SplitAfterN(l, " ", 2)
		hash := strings.TrimSpace(arr[0])
		hashes = append(hashes, hash)
		logMap[hash] = commit{
			commit: arr[1],
		}
	}
	for {
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
		options := []string{hashes[idxs[0]]}
		options = append(options, "--color=always")
		options = append(options, "--")
		content := git.Cmd("diff", options...)
		// if err != nil {
		// 	content = []byte("error reading conetent")
		// }
		pager.Show(string(content))
	}
}
