package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need only one argument")
		os.Exit(1)
	}

	words := strings.Fields(os.Args[1])

	var num int
	for _, word := range words {
		if len(word) > 0 {
			num++
		}
	}

	fmt.Println(num)
}
