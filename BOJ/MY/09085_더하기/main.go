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

func next() string {
	in.Scan()
	return in.Text()
}

func main() {
	Tc := nextInt()
	for i := 0; i < Tc; i++ {
		nextInt()
		var sum int = 0
		r := next()
		v := 0
		for _, j := range r {
			if j == ' ' {
				sum += v
				v = 0
				continue
			}
			v *= 10
			v += int(j) - '0'
		}
		sum += v
		fmt.Println(sum)
	}
}
