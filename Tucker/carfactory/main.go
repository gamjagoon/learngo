package main

import (
	"fmt"
	"strconv"
	"time"
)

type car struct {
	str string
}

func maketire(carChan chan car, outChan chan car) {
	for {
		car := <-carChan
		car.str += "Tire, "
		
		outChan <- car
	}
}

func makeengine(carChan chan car, outChan chan car)  {
	for {car := <- carChan
	car.str += "Engine, "

	outChan <- car}
}

func startwork(chan1 chan car)  {
	i := 1
	for {
		time.Sleep(1*time.Second)
		chan1 <- car{str:"Car"+strconv.Itoa(i)}
		i++
	}
}

/**.Golang multi thred
 * 1 Go Thread - lock unlock
 * 2 Go channel 
 * 3 Select 
 */

func main() {
	chan1 := make(chan car)
	chan2 := make(chan car)
	chan3 := make(chan car)
	/** chan1      chan2     chan3
	 *  main  | make tire | make engine
	 *  car+i->   tire    ->  engine
	 *                         |
	 *                         v
	 *                        print
	 */
	go startwork(chan1)
	go maketire(chan1, chan2)
	go makeengine(chan2, chan3)

	for {
		result := <- chan3
		fmt.Println(result.str)
	}

}