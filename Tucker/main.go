package main

import (
	"fmt"

	"github.com/gamjagoon/learngo/Tucker/datastruct"
)

func main() {
	fmt.Println("abcde = ", datastruct.Hash("abcde"))
	fmt.Println("abcde = ", datastruct.Hash("abcde"))

	mmap := datastruct.CreateMap()
	mmap.Add("aaa", "01088883334")
	mmap.Add("BBB", "01012312334")
	mmap.Add("ccc", "010f1231234")
	mmap.Add("lll", "01023123124")

	fmt.Println("aaa = ", mmap.Get("aaa"))
	fmt.Println("BBB = ", mmap.Get("BBB"))
	fmt.Println("ccc = ", mmap.Get("ccc"))
	fmt.Println("lll = ", mmap.Get("lll"))
}
