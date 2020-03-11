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
	for t := 1; main2(t); t++ {
	}
}

func main2(t int) bool {
	n, m := nextInt(), nextInt()
	if n == 0 && m == 0 {
		return false
	}
	a := make([]int, n+1)
	for i := range a {
		a[i] = i
	}
	b := make([]bool, n+1)
	b[0] = true
	var f func(x int) int
	f = func(x int) int {
		if a[x] != x {
			a[x] = f(a[x])
		}
		return a[x]
	}
	for ; m > 0; m-- {
		p, q := f(nextInt()), f(nextInt())
		a[q] = p
		if b[p] || b[q] {
			b[p] = true
			b[q] = true
		} else {
			b[q] = true
		}
	}
	r := 0
	fmt.Printf("Case %d: ", t)
	for _, v := range b {
		if !v {
			r++
		}
	}
	switch r {
	case 0:
		fmt.Printf("No trees.\n")
	case 1:
		fmt.Printf("There is one tree.\n")
	default:
		fmt.Printf("A forest of %d trees.\n", r)
	}
	return true
}
