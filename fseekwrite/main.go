package main

import (
	"io/ioutil"
	"os"
)

func main() {
	b := []byte("hello\ngo\n")
	filename := "file.data"
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b2, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	println("file is:")
	println(string(b2))

	ret, err := f.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	println("ret:", ret)

	b3 := []byte("bye\ngo\n")
	n, err := f.Write(b3)
	if err != nil {
		panic(err)
	}

	println("n:", n)

	err = f.Truncate(int64(n))
	if err != nil {
		panic(err)
	}

	ret, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	println("ret2:", ret)

	b4, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	println("b4:")
	println(string(b4))
}
