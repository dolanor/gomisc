package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	log.Fatal(run(os.Stdin))
}

func run(in io.Reader) error {
	times := []string{
		"2h30m",
		"1h30m",
		"10m",
		"21m",
		"1h",
		"17m",
		"56m",
		"1h07m",
		"1h34m",
		"1h05m",
		"1h53m",
		"48m",
		"12m",
		"1h26m",
		"2h36m",
		"1h34m",
		"1h10m",
		"1h08m",
		"3h40m",

		"22m",
		"6h25m",
		"13m",
		"1h34m",
		"1h45m",
		"2h11m",
		"1h11m",
		"38m",
		"27m",
		"24m",
		"1h",
		"8m",
		"2m",
		"35m",
	}

	var sum time.Duration
	for _, v := range times {
		d, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		sum += d
	}
	fmt.Println(sum)
	return nil
}
