package main

import (
	"strings"

	"github.com/chewxy/lingo/dep"
	"github.com/chewxy/lingo/lexer"
	"github.com/chewxy/lingo/pos"
)

func main() {
	inputString := `The cat sat on the mat`
	lx := lexer.New("dummy", strings.NewReader(inputString)) // lexer - required to break a sentence up into words.
	pt := pos.New(pos.WithModel(posModel))                   // POS Tagger - required to tag the words with a part of speech tag.
	dp := dep.New(depModel)                                  // Creates a new parser

	// set up a pipeline
	pt.Input = lx.Output
	dp.Input = pt.Output

	// run all
	go lx.Run()
	go pt.Run()
	go dp.Run()

	// wait to receive:
	for {
		select {
		case d := <-dp.Output:
			// do something
		case err := <-dp.Error:
			// handle error
		}
	}

}
