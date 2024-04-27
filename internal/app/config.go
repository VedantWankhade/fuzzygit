package app

import (
	"io"
	"log"
	"os"
)

var InfoLogger *log.Logger
var ErrorLogger *log.Logger

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
}
