package main

import "fmt"

type spoonofjam interface {
	String() string
}


// jam has getonespoon() method///
type jaminterface interface {
	getonespoon() spoonofjam
}

type bread struct {
	val string
}

type strawberryjam struct{

}

type orangejam struct {
}

func (s *spoonoforangejam) String () string{
	return "+ orange"
}

type spoonofstraberryjam struct{

}
type spoonoforangejam struct{

}

func (j *strawberryjam)getonespoon()spoonofjam {
	return &spoonofstraberryjam{}
}

func (j *orangejam)getonespoon()spoonofjam {
	return &spoonoforangejam{}
}

func (s* spoonofstraberryjam)String()string{
	return "+ strawberry"
}

func (b *bread)putjam(jam jaminterface) {
	spoon := jam.getonespoon()
	b.val += spoon.String()
}

func (b *bread) String()string {
	return "bread" +b.val
}

func main () {
	b := &bread{}
	//jam := &strawberryjam{}
	jam := &orangejam{}
	b.putjam(jam)
	fmt.Printf("%s", b)
}