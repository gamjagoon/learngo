package main

import "fmt"

func pop(c chan int)  {
	fmt.Println("Pop func")
	v := <-c
	fmt.Println(v)
}

func main()  {

	//*make channel 
	// make(chan int) : 0개 짜리 채널 따라서 아무것도 들어올수가 없다
	// make(chan int, size) : size개 짜리 채널 한번에 size 개 까지 처리가능
	c:= make(chan int)

	go pop(c)
	c <-10

	fmt.Println("end of program")
}