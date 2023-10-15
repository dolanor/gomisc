package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	a := os.Getenv("LD_TRACE_LOADED_OBJECT")
	if a != "" {
		log.Println("you're done goofed")
		return
	} else {
		fmt.Println("everything's fine")
	}

}
