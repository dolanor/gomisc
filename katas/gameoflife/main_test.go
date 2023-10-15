package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestGameOfLife(t *testing.T) {
	cases := map[string]struct {
		in  []string
		exp []string
	}{
		"1 dead => 1 dead": {
			in: []string{
				".",
			},
			exp: []string{
				".",
			},
		},

		"1 alive => 1 dead": {
			in: []string{
				"*",
			},
			exp: []string{
				".",
			},
		},

		"2 alive => dead": {
			in: []string{
				"**",
			},
			exp: []string{
				"..",
			},
		},

		"2 alive - => dead": {
			in: []string{
				"**.",
			},
			exp: []string{
				"...",
			},
		},

		//		"2 alive / => dead": {
		//			in: []string{
		//				"*.",
		//				".*",
		//			},
		//			exp: []string{
		//				"..",
		//				"..",
		//			},
		//		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := GameOfLife(c.in)
			is := is.New(t)
			is.Equal(got, c.exp)
		})
	}
}
