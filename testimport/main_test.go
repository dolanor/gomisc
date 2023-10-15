package main

import (
	"testing"

	"github.com/dolanor/gomisc/testimport/pkg"
	"github.com/dolanor/gomisc/testimport/pkg/testdata"
)

func TestPkgA(t *testing.T) {
	a := pkg.A{}
	aa := testdata.FakeA()
	_, _ = a, aa
}
