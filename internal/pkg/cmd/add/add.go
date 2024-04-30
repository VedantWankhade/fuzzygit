package add

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/config"
	"github.com/vedantwankhade/fuzzygit/internal/pkg/cmd/git"
)

func Invoke(flags []string) {
	flags = append(flags, "--exclude-standard", "-o", "-m")
	res := git.Cmd("ls-files", flags...)
	fileNames := strings.Split(res, "\n")
	f, err := fzf.New(fzf.WithNoLimit(true))
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(
		fileNames,
		func(i int) string { return fileNames[i] },
		fzf.WithPreviewWindow(func(i, width, height int) string {
			var options []string
			options = append(options, "--color=always")
			options = append(options, "--")
			options = append(options, "{-1}")
			options = append(options, fileNames[i])
			diff := git.Cmd("diff", options...)
			if diff == "" {
				fmt.Println("reading")
				config.InfoLogger.Println("reading file", fileNames[i])
				content, err := os.ReadFile(fileNames[i])
				if err != nil {
					config.ErrorLogger.Println("error reading", fileNames[i])
					fmt.Println("error reading", fileNames[i])
				}
				diff = string(content)
			}
			return diff
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(fileNames[i])
		_ = git.Cmd("add", fileNames[i])
	}
}
