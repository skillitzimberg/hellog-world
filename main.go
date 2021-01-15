package main

import (
	"os"
)

func main() {
	nf, err := os.Open("logs.txt")
	defer nf.Close()
	if err != nil {
		nf.Write([]byte("File creation failed."))
	}
	nf.Write([]byte("Hello, World!"))
}
