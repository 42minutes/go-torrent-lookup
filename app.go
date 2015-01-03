package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	term := "Big Bang Theory"

	resp, err := http.Get("https://torrentz.eu/search?q=" + strings.Replace(term, " ", "+", -1))
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(body)
		}

	} else {
		fmt.Println(err)
	}
}
