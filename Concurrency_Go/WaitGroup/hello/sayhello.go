package main

import (
	"fmt"
	"sync"
)


func main() {
  sayhello := func (ag *sync.WaitGroup, cnt int) {
    defer ag.Done()
    fmt.Printf("hello %v \n",cnt)
  }
  fmt.Print("start\n")
  var wg sync.WaitGroup
  wg.Add(10)
  for i := 0; i < 10; i++ {
    go sayhello(&wg,i)
  }
  wg.Wait()
  fmt.Print("end\n")
}