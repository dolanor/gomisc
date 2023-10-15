// +build gofuzzbeta

package pkg_test

import(
"testing"
pkg "github.com/dolanor/gomisc/fuzz"
)

func FuzzYolo2(f *testing.F) {
	f.Add("mdr")
	f.Add("MDr")
	f.Add("")
	f.Fuzz(func(t *testing.T, queryStr string) {
		ret := pkg.Yolo(queryStr)
		if ret == "YOLO" {
			t.Fatal("should not be empty")
		}

	})
}
