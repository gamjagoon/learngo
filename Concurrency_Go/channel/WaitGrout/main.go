package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
  begin := make(chan interface{})
  var wg sync.WaitGroup
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func (i int)  {
      defer wg.Done()
      <- begin // 계속진행해도 된다고 할 때까지 고루틴은 여기서 대기한다.
      fmt.Printf("%v has begun\n", i)
    }(i)
  }

  fmt.Println("Unblocking goruntines...")
  time.Sleep(2*time.Second)
  close(begin) // 닫는순간 대기상태에서벗어난다.
  wg.Wait()
}