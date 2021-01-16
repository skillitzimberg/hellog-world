package main

import (
	"log"
	"os"
)

func main() {
	lf, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer lf.Close()
	if err != nil {
		log.Printf("OpenFile failed:%s", err)
	}
	log.SetOutput(lf)
	log.Println("Hello, World!")
}
