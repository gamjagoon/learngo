package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var in = bufio.NewScanner(os.Stdin)

var n, r, c, s int

func main() {
	defer bufout.Flush()
	for {
		fmt.Fscan(bufin, &n, &r, &c)
		bufin.ReadBytes('\n')
		if n == 0 && r == 0 && c == 0 {
			break
		}
		s = (n+r+c)%25 + 1
		S := byte(s)
		str, _ := bufin.ReadBytes('\n')
		for _, i := range str {
			if 'a' <= i && i <= 'z' {
				i -= S
				if i < 'a' {
					i += 26
				}

			}
			bufout.WriteByte(i)
		}
	}
}
