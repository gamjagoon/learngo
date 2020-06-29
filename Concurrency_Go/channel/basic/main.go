package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
  // var dataStreamRead <-chan interface{}
  // dataStreamRead := make(<-chan interface{})

  // var dataStreamWrite chan<- interface{}
  // dataStreamWrite := make(chan<- interface{})

  stringStream := make(chan string)
  go func() {
    for i := 0; i < 3; i++ {
      time.Sleep(1*time.Second)
      fmt.Printf("%d sleep\n",i)
    }
    stringStream <- "Hello channels!"  
  }()

  str, ok := <-stringStream // 여기서 string이 나오기전까지 대기한다 
  if ok == false {
    log.Fatalln("not string in stringStream")
  }
  fmt.Println(str)

}