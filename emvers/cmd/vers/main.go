package main

import (
	_ "embed"
	"fmt"
	"runtime/debug"
)

//go:generate sh -c "git rev-parse --short HEAD >> version"

//go:embed version
var version string

func main() {
	print(version)
	fmt.Println(debug.ReadBuildInfo())
}
