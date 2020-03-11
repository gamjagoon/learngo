package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func read_int() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func read_stream() []byte {
	in.Scan()
	return in.Bytes()
}


func main() {
	N := read_int()
	a := make([]byte, 80)
	for i := 0; i < N; i++ {
		copy(a, read_stream())
		var (
			cal int = 0
			flag bool = false
		)
		for i,v := renge(a){
			if flag {

			}
		}
		fmt.Printf("%d\n", cal)
	}
}
