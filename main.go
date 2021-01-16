package main

import (
	"io"
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func init() {
	lf, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer lf.Close()
	if err != nil {
		log.Printf("Opening log file failed:%s", err)
	}
	infoLog = newLogger(lf, "INFO: ")
	infoLog.Println("Hellog, World!")
	errorLog = newLogger(lf, "ERROR: ")

}

func main() {
	cf, err := os.Create("output.txt")
	if err != nil {
		errorLog.Printf("Creating output file failed: %s", err)
	}

	cf.Write([]byte("New results: . . ."))
}

func newLogger(w io.Writer, p string) *log.Logger {
	return log.New(w, p, log.Ldate|log.Ltime|log.Lshortfile)
}
