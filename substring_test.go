// Tests detecting a substring in a string
// go test -bench . substring_test.go
package main

import (
	"regexp"
	"strings"
	"testing"
)

const (
	str    = "something"
	substr = "some"
)

var (
	re = regexp.MustCompile(substr)
)

func BenchmarkStringsContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Contains(str, substr)
	}
}

func BenchmarkStringsIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Index(str, substr)
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Split(str, substr)
	}
}

func BenchmarkStringsCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Count(str, substr)
	}
}

func BenchmarkStringsRegExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = regexp.MatchString(substr, str)
	}
}

func BenchmarkStringsRegExpCompiled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = re.MatchString(str)
	}
}
