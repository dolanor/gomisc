package main

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

var tmplText = `@font-face {
			font-family: 'ChronicleDisp-XLight';
			src: url('{{buildAssetsPath "ChronicleDisp-XLight.otf"}}');
		}`

var (
	outputDir = `C:\Windows\System32`
)

func main() {
	assetsPath := filepath.Join(outputDir, "assets")

	t, err := template.New("templ1").Funcs(template.FuncMap{
		"buildAssetsPath": func(filename string) template.URL {
			return template.URL(filepath.Join(assetsPath, filename))
		},
	}).Parse(tmplText)
	if err != nil {
		fmt.Printf("Template did not create for file %s. Err: %v\n", "<none>", err)
		return
	}

	var tplBuffer bytes.Buffer
	err = t.Execute(&tplBuffer, struct {
		Lines     []string
		AssetPath template.HTMLAttr
	}{[]string{"lol"}, template.HTMLAttr(assetsPath)})
	if err != nil {
		fmt.Printf("Template execution failed for file %s. Err: %v\n", "<none>", err)
	}

	fmt.Printf(tplBuffer.String())

}
