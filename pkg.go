// Package awssize provides a way to convert between AWS instance sizes,
// in a way that is mostly meaningful in the context of billing and reserved instances.
package awssize

import (
	"strings"
)

// Size represents AWS instance class size, like “medium”, “large”, and so on.
// Instances of this type should only be created with [S] function.
// Sizes can be compared agains each other using regular comparison operators,
// use [Size.As] method to step down from a larger value to some multiple of a smaller
// Size.
type Size int

// As returns the multiplier you need to apply to dst to get one src.
// It panics if src is smaller than dst, or if src cannot be expressed
// as non-fractional multiple of dst. It is recommended to use the
// smallest instance size within the family as dst.
func (src Size) As(dst Size) int {
	if src < dst {
		panic("size is smaller than target")
	}
	if src%dst != 0 {
		panic("size cannot be expressed as non-fractional multiple of the target")
	}
	return int(src / dst)
}

// S takes either an instance class, like “db.r6g.large”, “r5.large”,
// or a size suffix, like “large”, “medium”, etc., and returns the Size
// that can be used to calculate the difference between classes of
// the same family.
func S(s string) Size {
	v, ok := nameToSize[s[strings.LastIndexByte(s, '.')+1:]]
	if !ok {
		panic("unsupported size class")
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
