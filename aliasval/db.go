package main

import "errors"

var aliases = map[string]Alias{
	"stringAlias1": {Registry: "my", Key: "stringAlias1", Value: AliasString{value: "my alias 1"}},
	"arrayAlias1":  {Registry: "my", Key: "arrayAlias1", Value: AliasArray{value: []string{"array alias 1.1"}}},
	"arrayAlias2":  {Registry: "my", Key: "arrayAlias2", Value: AliasArray{value: []string{"array alias 2.1", "array alias 2.2"}}},
}

func GetAlias(id string) (*Alias, error) {
	alias, ok := aliases[id]
	if !ok {
		return nil, errors.New("alias not found")
	}
	return &alias, nil
}
