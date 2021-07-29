package main

import (
	"fmt"
	"log"
)

func main() {
	//deffer
	deffer()
	//deffer LIFO
	defferLOFO()
	//argumnt evaluated at the time defer is execute,not at time of  called function execution
	attimedeferexe()
	// panic recover
	fmt.Println("start")
	PanicRecover()
	fmt.Println("end")

}
func deffer() {
	defer fmt.Println("deffer 1")
	fmt.Println("deffer 2")
	fmt.Println("deffer 3")

}
func defferLOFO() {

	defer fmt.Println("deffer 1")
	defer fmt.Println("deffer 2")
	defer fmt.Println("deffer 3")

}

//argumnt evaluated at the time defer is execute,not at time of  called function execution
func attimedeferexe() {
	x := 1
	defer fmt.Println(x)
	x = 2

}

//only useful in deferd function
func PanicRecover() {
	fmt.Println("all about panic,recover and deffer")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		fmt.Println("after recover")

	}()
	panic("something bad")
	//fmt.Println("outside anonymous func") // don't execute

}
