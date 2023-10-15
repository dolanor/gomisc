package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var emailAddr, hash string
	_, err := fmt.Sscanf(os.Args[1], "%s|{SHA512-CRYPT}%s", &emailAddr, &hash)
	if err != nil {
		log.Println(err)
	}
}
