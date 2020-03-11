package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func next() string {
	in.Scan()
	return in.Text()
}

func main() {
	s := next()
	a := make([]string, nextInt())
	for i := range a {
		a[i] = next()
	}
	if regexp.MustCompile(fmt.Sprintf("^(%s)+$", strings.Join(a, "|"))).MatchString(s) {
		os.Stdout.WriteString("1")
	} else {
		os.Stdout.WriteString("0")
	}
}
