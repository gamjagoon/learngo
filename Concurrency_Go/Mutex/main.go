package main

import (
	"fmt"
	"sync"
)

func main() {
  var count int
  var lock sync.Mutex
  increment := func ()  {
    lock.Lock()
    defer lock.Unlock()

    count++
    fmt.Printf("증가 : %d\n", count)
  }

  decrement := func ()  {
    lock.Lock()
    defer lock.Unlock()

    count--
    fmt.Printf("감소 : %d\n", count)
  }

  var wg sync.WaitGroup
  for i := 0; i <=  5; i++ {
    wg.Add(1)
    go func ()  {
      defer wg.Done()
      increment()
    }()
  }

  for i := 0; i <=  5; i++ {
    wg.Add(1)
    go func ()  {
      defer wg.Done()
      decrement()
    }()
  }

  wg.Wait()
  fmt.Println("end", count)
}