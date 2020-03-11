package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ y, x int }

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var in = bufio.NewScanner(os.Stdin)
var dir = [...]pos{{0, -1}, {0, 0}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}}
var maps [20][][]byte
var visited [20][8][8]bool

// N = 8 X 8 arr
var N int = 8

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

func outrange(x, y int) bool {
	return x < 0 || y < 0 || x >= 8 || y >= 8
}

func dfs(x, y, t int) int {
	if t > 18 {
		return 0
	}
	if x == 0 && y == 7 {
		return 1
	}
	res := 0
	var nx, ny int
	for i := range dir {
		nx = x + dir[i].x
		ny = y + dir[i].y
		if !outrange(nx, ny) && !visited[t+1][nx][ny] && maps[t][nx][ny] == '.' && maps[t+1][nx][ny] == '.' {
			visited[t+1][nx][ny] = true
			res |= dfs(nx, ny, t+1)
		}
	}
	return res
}

func main() {
	defer bufout.Flush()
	orimap := make([][]byte, 8)
	for i := range orimap {
		orimap[i] = make([]byte, 8)
		copy(orimap[i], next())
	}
	maps[0] = orimap
	for i := 1; i <= 19; i++ {
		tmpmap := make([][]byte, 8)
		for j := 0; j < i && j <= 7; j++ {
			tmpmap[j] = make([]byte, 8)
			for k := range tmpmap[j] {
				tmpmap[j][k] = '.'
			}
		}
		v := 0
		for j := i; j <= 7; j++ {
			tmpmap[j] = make([]byte, 8)
			copy(tmpmap[j], orimap[v])
			v++
		}
		maps[i] = tmpmap
	}
	visited[0][7][0] = true
	fmt.Print(dfs(7, 0, 0))
}
