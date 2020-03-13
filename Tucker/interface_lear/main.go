package main

import "fmt"

/**
 * ! interface
 * ? A Class
 * ? |-------b method
 * ? |-------c method
 * * B Class
 * * |-------d method
 * * |-------e method
 * 
 * ? interfase -> external B
 */
type jam interface {
	getonespoon() *spoonofstarwberryjam
}

type bread struct {
	val string
}
type strawbrerryham struct {

}
type spoonofstarwberryjam struct {
	
}


func (s *spoonofstarwberryjam)sstring() string{
	return "+ stawberry"
}

func (j *strawbrerryham)getonespoon() *spoonofstarwberryjam {
	return &spoonofstarwberryjam{}
}

func (b *bread) putjam(jam *strawbrerryham){
	spoon := jam.getonespoon()
	b.val += spoon.sstring()
}

func (b *bread)sstring() string {
	return "bread" + b.val
}

func main()  {
	b := &bread{}
	jam := &strawbrerryham{}
	b.putjam(jam)
	fmt.Println(b.val)
}