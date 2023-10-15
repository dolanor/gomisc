package pkg

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	var r io.Reader
	r = strings.NewReader("MDR=postgres://localhost:5432/db?sslmode=disable")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	r = bytes.NewReader(bytes.ReplaceAll(b, []byte("="), []byte(" ")))
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}

	fmt.Println(string(b))
}
