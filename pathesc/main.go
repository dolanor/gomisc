package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/"+"123 et 456", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(req.URL.Opaque)
	fmt.Println(req.URL.Path)
	fmt.Println(req.URL.RawPath)
	fmt.Println(req.URL.EscapedPath())

}
