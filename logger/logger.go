package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

func NewFileLogger(filepath string, prefix string) *log.Logger {
	f, err := os.OpenFile("log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger := log.New(f, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	return logger
}
func NewLogger(out io.Writer, prefix string) *log.Logger {
	logger := log.New(out, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	return logger
}
