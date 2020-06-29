package main

import (
	"fmt"
	"time"
)

func main() {
  start := time.Now()
  c := make(chan interface{})
  go func ()  {
    time.Sleep(5*time.Second)
    close(c)
  }()

  fmt.Println("Blocking on read...")
  loop:
  for {
    select {
    case <-c :
      fmt.Printf("Unblocked %v later.\n",time.Since(start))
      break loop
    default:
      fmt.Println("pass")
    }
    time.Sleep(1*time.Second)
  }
}