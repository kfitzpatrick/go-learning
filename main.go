package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf)
}
