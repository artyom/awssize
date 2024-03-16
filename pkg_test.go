package awssize

import (
	"fmt"
	"testing"
)

func TestS(t *testing.T) {
	for _, tc := range []struct {
		mul int
		x   string
		y   string
	}{
		{2, "cache.t3.small", "cache.t3.medium"},
		{8, "medium", "2xlarge"},
		{48 * 2, "large", "48xlarge"},
	} {
		if Size(tc.mul)*S(tc.x) != S(tc.y) {
			t.Fatalf("%d %s != %s, but expected to be equal", tc.mul, tc.x, tc.y)
		}
	}
}

func ExampleSize_As() {
	src := S("2xlarge")
	dst := S("medium")
	fmt.Printf("one %s equals %d %s\n", src, src.As(dst), dst)
	fmt.Printf("one %s equals %d %s\n", src, src.As(src), src)
	// Output:
	// one 2xlarge equals 8 medium
	// one 2xlarge equals 1 2xlarge
}
