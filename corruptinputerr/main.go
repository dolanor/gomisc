package main

import (
	"encoding/base64"
	"log"
)

func main() {
	str := "ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="
	//str := "ROAZBWtSacxXQrOe3FGAqJDyJjFe]PR5ce4TSIzmJ0Bc=o"
	_, err := base64.StdEncoding.DecodeString(str)
	if cie, ok := err.(base64.CorruptInputError); ok {
		i := int64(cie)
		rstr := []byte(str)
		und := string('̭')
		newd := []byte(und)

		log.Println("yop:", string(rstr[:i+1]))

		newStr := append(rstr[:i+1], newd...)
		log.Println("yop:", string(newStr))
		newStr = append(newStr, rstr[i+1:]...)
		log.Println("err:", err, ":", string(newStr))
	}
}
