package main

import (
	"bytes"
	"text/template"
)

func main() {
	type MDR struct {
		Name string
	}

	t, err := template.New("MyName").Parse("lol {{ .Name }}")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, MDR{Name: "mdrName"})
	if err != nil {
		panic(err)
	}

	//tree, err := parse.Parse("MyName", "lol {{ MDR }}", "{{", "}}", nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(tree)
}
