package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var tc, a, c, r, l int
var line = [10001]bool{false}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &tc)
	r = 1
	l = 10000
	for ; tc > 0; tc-- {
		fmt.Fscan(bufin, &a, &c)
		for i := a; i < c; i++ {
			line[i] = true
		}
	}
	s := 0
	for i := 1; i < 10000; i++ {
		if line[i] {
			s++
		}
	}
	fmt.Fprint(bufout, s)
}
