package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: wordcount PHRASE")
		os.Exit(1)
	}

	words := strings.Fields(os.Args[1])

	fmt.Println(len(words))
}
