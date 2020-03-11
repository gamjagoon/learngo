package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, r, c int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &r, &c)
	for i := 0; i < n; i++ {
		for j := 0; j < (n / 2); j++ {
			if (r%2 == 0 && c%2 == 0) || (r%2 != 0 && c%2 != 0) {
				if i%2 != 0 {
					fmt.Fprint(bufout, ".v")
				} else {
					fmt.Fprint(bufout, "v.")
				}
			} else {
				if i%2 == 0 {
					fmt.Fprint(bufout, ".v")
				} else {
					fmt.Fprint(bufout, "v.")
				}
			}
		}
		if n%2 != 0 {
			if (r%2 == 0 && c%2 == 0) || (r%2 != 0 && c%2 != 0) {
				if i%2 != 0 {
					fmt.Fprint(bufout, ".")
				} else {
					fmt.Fprint(bufout, "v")
				}
			} else {
				if i%2 == 0 {
					fmt.Fprint(bufout, ".")
				} else {
					fmt.Fprint(bufout, "v")
				}
			}

		}
		fmt.Fprint(bufout, "\n")
	}
}
