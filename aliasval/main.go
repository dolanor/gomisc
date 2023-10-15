package main

import (
	"fmt"
	"log"
)

func main() {
	// create new alias array value with data from DB
	av, err := NewAliasValue("array", []string{"yo ho"})
	if err != nil {
		panic(err)
	}
	fmt.Println(av)

	// create new alias array value with data from DB
	av, err = NewAliasValue("string", "yo ho")
	if err != nil {
		panic(err)
	}
	fmt.Println(av)

	// create a wrong alias value with data from DB
	av, err = NewAliasValue("what?", "yo ho")
	if err != nil {
		log.Println("wrong alias value:", err)
	}
	fmt.Println(av)

	a, err := GetAlias("stringAlias1")
	if err != nil {
		panic(err)
	}
	// Check 1 kind
	aa, ok := a.Value.(AliasString)
	if !ok {
		panic("oops")
	}
	fmt.Println(aa)

	for k := range aliases {
		fmt.Println("==", k)
		a, err := GetAlias(k)
		if err != nil {
			panic(err)
		}
		// Check multiple kind
		switch t := a.Value.(type) {
		case AliasArray:
			fmt.Println("array:", t.Value())
		case AliasString:
			fmt.Println("string:", t.Value())

		}
	}
}
