package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func main() {
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func run(in io.Reader, out io.Writer) error {
	csvr := csv.NewReader(in)
	records, err := csvr.ReadAll()
	if err != nil {
		return err
	}

	var dates []time.Time
	for _, v := range records {
		rfcDate := fmt.Sprintf("%sT%sZ", v[0], v[1])
		fmt.Fprintln(os.Stderr, rfcDate)
		date, err := time.Parse(time.RFC3339, rfcDate)
		if err != nil {
			return err
		}
		dates = append(dates, date)
	}

	var start, curr time.Time
	sessions := map[string]time.Duration{}

	for _, d := range dates {
		if curr.Sub(d) > time.Minute*30 {
			session := start.Sub(curr)
			fmt.Println(session)
			if session < 0 && session < -24*time.Hour {
				start = d
				curr = d
				continue
			}
			if session < 0 {
				session *= -1
			}
			sessionRange := start.Format("2006-01-02 15:04-") + curr.Format("15:04")
			sessions[sessionRange] = session
			start = d
		}
		curr = d
	}
	//sessions = sessions[1:]
	fmt.Fprintln(os.Stderr, sessions)

	var sum time.Duration
	var list []string
	for k, d := range sessions {
		//	fmt.Fprintln(out, k, d)
		list = append(list, fmt.Sprint(k, " ", d))

		sum += d
	}

	ss := sort.StringSlice(list)
	ss.Sort()
	sort.Sort(sort.Reverse(ss))
	for _, v := range ss {
		fmt.Println(v)
	}
	fmt.Fprintln(os.Stderr, sum)
	return nil
}
