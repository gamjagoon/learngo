package main

import (
	"fmt"
	"strconv"
	"time"
)

type car struct {
	str string
}

type plane struct {
	str string
}

func maketire(carChan chan car, outChan chan car,planeChan,outplaneChan chan plane) {
	for {
		select {
		case car:= <-carChan:
			car.str += "Tire, "
			outChan <- car
		case plane:= <-planeChan:
			plane.str += "Wheel, "
			outplaneChan <- plane
		}
	}
}

func makeengine(carChan chan car, outChan chan car,planeChan,outplaneChan chan plane)  {
	for {
		select {
		case car:= <-carChan:
			car.str += "carEngine, "
			outChan <- car
		case plane:= <-planeChan:
			plane.str += "planeEngine, "
			outplaneChan <- plane
		}
	}
}

func starcartwork(chancar chan car)  {
	i := 1
	for {
		time.Sleep(1*time.Second)
		chancar <- car{str:"Car"+strconv.Itoa(i)}
		i++
	}
}


func starplanetwork(chanpl chan plane)  {
	i := 1
	for {
		time.Sleep(1*time.Second)
		chanpl <-plane{str:"Plane"+strconv.Itoa(i)}
		i++
	}
}

/**.Golang multi thred
 * 1 Go Thread - lock unlock
 * 2 Go channel 
 * 3 Select 
 */

func main() {
	carchan1 := make(chan car)
	carchan2 := make(chan car)
	carchan3 := make(chan car)

	planechan1 := make(chan plane)
	planechan2 := make(chan plane)
	planechan3 := make(chan plane)
	/** chan1      chan2     chan3
	 *  main  | make tire | make engine
	 *  car+i->   tire    ->  engine
	 *                         |
	 *                         v
	 *                        print
	 */
	go starcartwork(carchan1)
	go starplanetwork(planechan1)
	go maketire(carchan1, carchan2,planechan1,planechan2)
	go makeengine(carchan2, carchan3,planechan2,planechan3)

	for {
		select {
		case resultcar := <- carchan3 :
			fmt.Println(resultcar.str)
		case resultplane := <- planechan3:
			fmt.Println(resultplane.str)
		default:
			fmt.Print("Not input")
		}
	}

}