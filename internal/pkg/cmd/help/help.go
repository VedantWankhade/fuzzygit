package help

import (
	"log"

	"github.com/koki-develop/go-fzf"
	"github.com/vedantwankhade/fuzzygit/internal/config"
)

func Invoke(flags []string) {
	var helpItems []string
	for k := range config.HelpContent {
		helpItems = append(helpItems, k)
	}
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Find(
		helpItems,
		func(i int) string { return helpItems[i] },
		fzf.WithPreviewWindow(func(i, width, height int) string {
			return config.HelpContent[helpItems[i]]
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
