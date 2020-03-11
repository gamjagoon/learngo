package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

//M ...
const M int64 = 1234567891

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func next() []byte {
	in.Scan()
	return in.Bytes()
}

func hashcode(str []byte, len int) int64 {
	if len == 1 {
		return int64(str[0] - 96)
	}
	var res int64 = int64(str[0] - 96)
	var com int64 = 1
	for i := 1; i < len; i++ {
		com = (com * 31) % M
		res += (int64(str[i]-96) * com) % M
	}
	return res % M
}

func main() {
	N := nextInt()
	a := make([]byte, N)
	copy(a, next())
	fmt.Println(hashcode(a, N))
	return
}
