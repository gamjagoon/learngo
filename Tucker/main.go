package main

import (
	"fmt"

	"github.com/gamjagoon/learngo/Tucker/datastruct"
)

func main() {
	h := &datastruct.Heap{}
	for i := 1; i < 10; i++ {
		h.Push(i)
	}
	h.Print()
	for i := 1; i < 10; i++ {
		fmt.Println(h.Pop())
	}
}
