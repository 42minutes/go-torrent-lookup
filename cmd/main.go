package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tl "github.com/42minutes/go-torrentlookup"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Missing query arg")
	}

	n, i, err := tl.ProviderTPB.Search(strings.Join(os.Args[1:], " "))
	fmt.Println(n, i, tl.CreateFakeMagnet(i), err)
}
