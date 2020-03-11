package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	in.Split(bufio.ScanWords)
	n := nextInt()
	r := make([]int, n)
	r[0] = 1
	for i := range r {
		x := nextInt()
		if r[i] == 0 {
			continue
		}
		for j := 1; j <= x; j++ {
			if k := i + j; k < n && (r[k] == 0 || r[k] > r[i]+1) {
				r[k] = r[i] + 1
			}
		}
	}
	if r[n-1] == 0 {
		fmt.Printf("-1")
	} else {
		fmt.Print(r[n-1] - 1)
	}
}
