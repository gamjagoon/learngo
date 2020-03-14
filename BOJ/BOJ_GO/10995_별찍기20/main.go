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

func printStars(N int) {
	for i := 0; i < N-1; i++ {
		fmt.Print("* ")
	}
	fmt.Println("*")
}

func main() {
	N := nextInt()
	if N == 1 {
		fmt.Println("*")
	} else {
		for i := 0; i < N; i++ {
			if i&1 == 0 {
				printStars(N)
			} else {
				fmt.Print(" ")
				printStars(N)
			}
		}
	}
}