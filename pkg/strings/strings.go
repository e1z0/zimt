package strings

import "strings"

// MaskSym keep the symbol to use on masking
const MaskSym = "*"

// Mask replaces the most of string characters with MaskSym
func Mask(key string) string {
	l := len(key)

	if l == 0 {
		return key
	}
	if l == 1 {
		return MaskSym
	}
	if l < 5 {
		return key[0:1] + strings.Repeat(MaskSym, l-1)
	}
	if l < 7 {
		return key[0:1] + strings.Repeat(MaskSym, l-2) + key[l-1:]
	}
	if l < 9 {
		return key[0:2] + strings.Repeat(MaskSym, l-3) + key[l-1:]
	}
	return key[0:2] + strings.Repeat(MaskSym, 10) + key[l-2:]
}
