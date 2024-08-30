package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := os.Args[1]
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer r.Body.Close()

	file, err := os.Create("page.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

}
