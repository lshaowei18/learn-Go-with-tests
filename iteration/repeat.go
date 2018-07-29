package iteration

import "strings"

func Repeat(s string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}

func ExampleRepeat(s string, n int) string {
	repeated := ""
	for i := 0; i < n; i++ {
		repeated += s
	}
	return repeated
}

func StringsRepeat(s string, n int) string {
	return strings.Repeat(s, n)
}
