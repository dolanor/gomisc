package main

import "log"

type State int

func (s State) String() string {
	return stateMap[s]
}

const (
	bad = iota
	good
	maybe
)

var stateMap = map[State]string{
	bad:   "bad",
	good:  "good",
	maybe: "maybe",
}

func run(st State) {
	log.Println(st)
}

func main() {
	run(good)
}
