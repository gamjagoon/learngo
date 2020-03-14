package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var a, b int
var arr = make([]int, 2000)

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &a, &b)
	i := 1
	t := 1
	for ; i < 70; i++ {
		for j := 0; j < i; j++ {
			arr[t] = i
			t++

		}
		if t > b {
			break
		}
	}
	var res = 0
	for i = a; i <= b; i++ {
		res += arr[i]
	}
	fmt.Fprintln(bufout, res)
}
