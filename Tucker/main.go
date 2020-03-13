package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		go func1()
		fmt.Println("main", i)
	}
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("func 1 : ", i)
	}
}
