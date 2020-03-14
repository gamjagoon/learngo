package main

import (
	"fmt"
	"strconv"
)

type structa struct {
	val string
}
func (s * structa)String() string{
	return "Val :"+s.val
}

type structb struct {
	val int
}

func (s *structb) String()string{
	return "structb: "+strconv.Itoa(s.val)
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}
func main() {
	a := &structa{val:"AAA"}
	b := &structb{val:1234}
	Println(a)
	Println(b)
}