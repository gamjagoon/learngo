package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, p int
var nscore uint
var arr = [51]uint{}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &nscore, &p)
	if n == 0 {
		fmt.Println(1)
		os.Exit(0)
	}
	var rank, sp int = 1, 1
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &arr[i])
		if arr[i] > nscore {
			rank++
		}
		if arr[i] >= nscore {
			sp++
		}
	}
	if rank > p {
		fmt.Println(-1)
	} else if sp > p {
		fmt.Println(-1)
	} else {
		fmt.Println(rank)
	}
}
