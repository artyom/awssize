// Package awssize provides utilities for comparing and converting between
// different AWS EC2 instance sizes within the same family. It allows representing
// sizes like "large", "xlarge", etc. as values that can be easily compared and
// used to calculate size differences.
package awssize

import (
	"strings"
)

// Size represents an AWS EC2 instance size like "medium", "large", "xlarge", etc.
// Size values can be compared using regular comparison operators.
// Use the [S] function to create Size values from either full instance classes
// like "r5.large" or just the size suffix like "large".
// The zero value is not a valid Size.
type Size int

// As returns the multiplier needed to convert from the source Size to the
// destination Size dst. For example:
//
//	xlarge := S("xlarge")
//	large := S("large")
//	fmt.Println(xlarge.As(large)) // Prints 2
//
// As panics if the source Size is smaller than dst, or if it cannot be expressed
// as a whole multiple of dst. For best results, use the smallest size in the
// instance family as the destination, e.g. S("large").
func (src Size) As(dst Size) int {
	if src < dst {
		panic("size is smaller than target")
	}
	if src%dst != 0 {
		panic("size cannot be expressed as non-fractional multiple of the target")
	}
	return int(src / dst)
}

// S parses an instance size string and returns the corresponding Size value.
// s can be either a full instance class like "db.r6g.large", "r5.large", etc.
// or just the size suffix like "large", "xlarge", "2xlarge", etc.
// It panics if the input string does not contain a valid size suffix.
func S(s string) Size {
	v, ok := nameToSize[s[strings.LastIndexByte(s, '.')+1:]]
	if !ok {
		panic("unsupported size class: " + s)
	}
	return v
}

var nameToSize map[string]Size

func init() {
	nameToSize = make(map[string]Size)
	for k, v := range _Size_map {
		nameToSize[v] = k
	}
}

const (
	sizeNano     Size = 1               // nano
	sizeMicro    Size = 2 * sizeNano    // micro
	sizeSmall    Size = 2 * sizeMicro   // small
	sizeMedium   Size = 2 * sizeSmall   // medium
	sizeLarge    Size = 2 * sizeMedium  // large
	sizeXlarge   Size = 2 * sizeLarge   // xlarge
	size2xLarge  Size = 2 * sizeXlarge  // 2xlarge
	size3xLarge  Size = 3 * sizeXlarge  // 3xlarge
	size4xLarge  Size = 4 * sizeXlarge  // 4xlarge
	size8xLarge  Size = 8 * sizeXlarge  // 8xlarge
	size9xLarge  Size = 9 * sizeXlarge  // 9xlarge
	size10xLarge Size = 10 * sizeXlarge // 10xlarge
	size12xLarge Size = 12 * sizeXlarge // 12xlarge
	size16xLarge Size = 16 * sizeXlarge // 16xlarge
	size18xLarge Size = 18 * sizeXlarge // 18xlarge
	size24xLarge Size = 24 * sizeXlarge // 24xlarge
	size32xLarge Size = 32 * sizeXlarge // 32xlarge
	size48xLarge Size = 48 * sizeXlarge // 48xlarge
)

//go:generate stringer -type=Size -linecomment
