package config

import (
	"io"
	"log"
	"os"
)

var InfoLogger *log.Logger
var ErrorLogger *log.Logger
var HelpContent map[string]string

const logFilePath string = "/tmp/fuzzygit/fuzzygit.log"

func init() {
	logWritter := io.Discard
	tmpDir := os.TempDir()
	err := os.MkdirAll(tmpDir+"/fuzzygit", os.ModePerm)
	if err == nil {
		logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err == nil {
			logWritter = logFile
		}
	}
	InfoLogger = log.New(logWritter, "INFO", log.Ltime|log.Ldate|log.Lshortfile)
	ErrorLogger = log.New(logWritter, "ERROR", log.Ltime|log.Ldate|log.Lshortfile)
	HelpContent = map[string]string{
		"Print this help menu":       "fuzzygit help",
		"Show diff of tracked files": "fuzzygit diff",
		"Show diff of staged files":  "fuzzygit diff --staged",
		"Show diff of two commits":   "fuzzygit diff -c",
		"Stage changes(files)":       "fuzzygit add",
		"Checkout local branch":      "fuzzygit checkout",
		"Checkout remote branch":     "fuzzygit checkout -r",
		"Checkout a tag (creates a local branch with name 'branch-<tagname>')":       "fuzzygit checkout -t",
		"Checkout a commit (creates a local branch with name 'branch-<commithash>')": "fuzzygit checkout -c",
		"Show commit logs/history":                        "fuzzygit log",
		"Show commits that whose diff has a given string": "fuzzygit log -S <string>",
		"Rename a local branch":                           "fuzzygit rename",
		"Unstage staged files":                            "fuzzygit unstage",
	}
}
