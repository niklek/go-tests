// Detecting a substring in a string code examples
package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	str    = "something"
	substr = "some"
)

var (
	re = regexp.MustCompile(substr)
)

func main() {
	// Contains
	res := strings.Contains(str, substr)
	fmt.Println(res) // true

	// Index: check the index of the first instance of substr in str, or -1 if substr is not present
	i := strings.Index(str, substr)
	fmt.Println(i) // 0

	// Split by substr and check len of the slice, or length is 1 if substr is not present
	ss := strings.Split(str, substr)
	fmt.Println(len(ss)) // 2

	// Check number of non-overlapping instances of substr in str
	c := strings.Count(str, substr)
	fmt.Println(c) // 1

	// RegExp
	matched, _ := regexp.MatchString(substr, str)
	fmt.Println(matched) // true

	// Compiled RegExp
	res = re.MatchString(str)
	fmt.Println(res) // true
}
