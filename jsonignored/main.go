package main

import "encoding/json"

func main() {
	type A struct {
		ID     int
		DBOnly string
	}

	b, err := json.Marshal(&A{1, "lol"})
	if err != nil {
		println(err.Error())
	}
	println(string(b))
}
