package main

import (
	"fmt"
	"sync"
)

func main() {
  c := sync.NewCond(&sync.Mutex{})
  queue := make([]interface{},0,10)

  removeFromQueue := func (q *[]interface{}) {
    c.L.Lock()
    *q =(*q)[1:]
    fmt.Println("Remove to queue")
    fmt.Println(*q)
    c.L.Unlock()
    c.Signal()
  }

  for i := 0; i < 10; i++ {
    c.L.Lock()
    for len(queue) == 2{
      c.Wait()
    }
    queue = append(queue,i)
    fmt.Println("Adding to queue")
    fmt.Println(queue)
    go removeFromQueue(&queue)
    c.L.Unlock()
  }
  fmt.Println(queue)

}