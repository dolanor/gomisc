package main

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
)

//go:embed assets
var assets embed.FS

func main() {
	entries, err := fs.ReadDir(assets, "assets")
	fmt.Println(err)
	for _, v := range entries {
		if v.IsDir() {
			continue
		}
		data, err := assets.ReadFile(filepath.Join("assets", v.Name()))
		fmt.Println(string(data), err)
	}
}
