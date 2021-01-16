package main

import (
	"log"
	"os"
)

func main() {
	lf, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer lf.Close()
	if err != nil {
		log.Printf("Opening log file failed:%s", err)
	}
	log.SetOutput(lf)
	log.Println("Hellog, World!")

	cf, err := os.Create("output.txt")
	if err != nil {
		log.Printf("Creating output file failed: %s", err)
	}

	cf.Write([]byte("Here are your results: . . ."))
}
