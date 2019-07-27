package strings

import "testing"

func TestMask(t *testing.T) {
	samples := []struct {
		in  string
		out string
	}{
		{
			in:  "",
			out: "",
		}, {
			in:  "a",
			out: "*",
		}, {
			in:  "ab",
			out: "a*",
		}, {
			in:  "abc",
			out: "a**",
		}, {
			in:  "abcd",
			out: "a***",
		}, {
			in:  "abcde",
			out: "a***e",
		}, {
			in:  "abcdef",
			out: "a****f",
		}, {
			in:  "abcdefg",
			out: "ab****g",
		}, {
			in:  "abcdefgh",
			out: "ab*****h",
		}, {
			in:  "abcdefghi",
			out: "ab**********hi",
		}, {
			in:  "abcdefghij",
			out: "ab**********ij",
		}, {
			in:  "abcdefghijk",
			out: "ab**********jk",
		}, {
			in:  "abcdefghijkl",
			out: "ab**********kl",
		}, {
			in:  "abcdefghijklm",
			out: "ab**********lm",
		},
	}

	for _, s := range samples {
		if res := Mask(s.in); res != s.out {
			t.Errorf("Expected %s, got %s", s.out, res)
		}
	}
}
