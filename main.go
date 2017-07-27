package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	writtenBytes, errCopy := io.Copy(os.Stdout, f)

	if errCopy != nil {
		panic(errCopy)
	}
	fmt.Printf("\n")
	fmt.Printf("%d bytes written", writtenBytes)
}
