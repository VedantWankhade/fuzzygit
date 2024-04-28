package diff

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/config"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
)

func Invoke(flags []string) {
	flags = append(flags, "--name-only")
	res := git.Cmd("diff", flags...)
	fileNames := strings.Split(res, "\n")
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
}
