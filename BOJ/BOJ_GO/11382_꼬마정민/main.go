package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
	in.Scan()
	var r int64 = 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int64(c - '0')
	}
	return r
}

func main() {
	var sum int64 = 0
	in.Split(bufio.ScanWords)
	for i := 0; i < 3; i++ {
		sum += nextInt()
	}
	fmt.Println(sum)
}
