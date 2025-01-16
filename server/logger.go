package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Starts logging to STDOUT and .makao.log file to ./logs directory.
// Returns cleanup function which closes the file.
func setupLogger() (closeCallback func()) {
	logPath := filepath.Join(".", "logs")
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		log.Fatalf("error opening ./logs dir.: %v", err)
	}

	fileName := fmt.Sprintf("%d.makao.log", time.Now().Unix())
	logFilePath := filepath.Join(logPath, fileName)
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	return func() {
		f.Close()
	}
}
