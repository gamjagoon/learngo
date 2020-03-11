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

func next() []byte {
	in.Scan()
	return in.Bytes()
}

func main() {
	in.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	a := make([][]byte, n)
	for i := range a {
		a[i] = make([]byte, m)
		copy(a[i], next())
	}
	r := 32
	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			ra, rb := 0, 0
			for p := 0; p < 8; p++ {
				for q := 0; q < 8; q++ {
					if "BW"[(p+q)&1] == a[i+p][j+q] {
						ra++
					} else {
						rb++
					}
				}
			}
			if r > ra {
				r = ra
			}
			if r > rb {
				r = rb
			}
		}
	}
	fmt.Println(r)
}
