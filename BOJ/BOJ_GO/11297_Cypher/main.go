package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, r, c int

func carrot(odd bool) {
	if odd {
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				for j := 0; j < n; j++ {
					if j&1 == 0 {
						fmt.Fprint(bufout, "v")
					} else {
						fmt.Fprint(bufout, ".")
					}
				}
			} else {
				for j := 0; j < n; j++ {
					if j&1 == 1 {
						fmt.Fprint(bufout, "v")
					} else {
						fmt.Fprint(bufout, ".")
					}
				}
			}
			fmt.Fprint(bufout, "\n")
		}
	} else {
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				for j := 0; j < n; j++ {
					if j&1 == 1 {
						fmt.Fprint(bufout, "v")
					} else {
						fmt.Fprint(bufout, ".")
					}
				}
			} else {
				for j := 0; j < n; j++ {
					if j&1 == 0 {
						fmt.Fprint(bufout, "v")
					} else {
						fmt.Fprint(bufout, ".")
					}
				}
			}
			fmt.Fprint(bufout, "\n")
		}
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &r, &c)
	var firstodd bool
	if r&1 == 1 {
		if c&1 == 1 {
			firstodd = true
		} else {
			firstodd = false
		}
	} else {
		if c&1 == 1 {
			firstodd = false
		} else {
			firstodd = true
		}
	}
	carrot(firstodd)
}
