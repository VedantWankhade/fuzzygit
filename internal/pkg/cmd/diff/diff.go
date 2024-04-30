package diff

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/config"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/pager"
)

func Invoke(flags []string) {
	if slices.Contains(flags, "-c") {
		type commit struct {
			commit string
		}
		logMap := make(map[string]commit)
		var hashes []string
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
		f, err := fzf.New(fzf.WithLimit(2))
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
		var options []string
		for _, i := range idxs {
			options = append(options, hashes[i])
		}
		options = append(options, "--color=always")
		options = append(options, "--")
		content := git.Cmd("diff", options...)
		// if err != nil {
		// 	content = []byte("error reading conetent")
		// }
		pager.Show(string(content))
	} else {
		flags = append(flags, "--name-only")
		res := git.Cmd("diff", flags...)
		fileNames := strings.Split(res, "\n")
		for {
			f, err := fzf.New()
			if err != nil {
				log.Fatal(err)
			}
			idxs, err := f.Find(
				fileNames,
				func(i int) string { return fileNames[i] },
				fzf.WithPreviewWindow(func(i, width, height int) string {
					var options []string
					if slices.Contains(flags, "--staged") {
						options = append(options, "--staged")
					}
					options = append(options, "--color=always")
					options = append(options, "--")
					options = append(options, "{-1}")
					options = append(options, fileNames[i])
					diff := git.Cmd("diff", options...)
					config.InfoLogger.Println(string(diff))
					return fmt.Sprintf("%v", string(diff))
				}),
			)
			if err != nil {
				log.Fatal(err)
			}

			for _, i := range idxs {
				fmt.Println(fileNames[i])
			}
			content, err := os.ReadFile(fileNames[idxs[0]])
			if err != nil {
				content = []byte("error reading conetent")
			}
			pager.Show(string(content))
		}
	}
}
