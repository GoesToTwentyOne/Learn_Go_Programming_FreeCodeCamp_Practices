package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 21

	}()
	func() {
		fmt.Println(<-ch)

	}()

	//multipleDataSendReceived
	multipleDataSendReceived()
	//bedirectionalUnbuffer
	bedirectionalUnbuffer()
	//bedirectionalbuffer
	bedirectionalbuffer()
	fmt.Println("all exit")
}
func multipleDataSendReceived() {
	var wg sync.WaitGroup
	wg.Add(2)

	chm := make(chan int)
	go func(c chan int) {

		fmt.Println(<-chm)
		chm <- 21
		wg.Done()

	}(chm)
	go func(c chan int) {
		chm <- 41
		fmt.Println(<-chm)
		wg.Done()

	}(chm)
	wg.Wait()

}
func bedirectionalUnbuffer() {
	chm := make(chan int)
	go func(c <-chan int) {

		fmt.Println(<-chm)

	}(chm)
	func(c chan<- int) {
		chm <- 41

	}(chm)
}
func bedirectionalbuffer() {
	chm := make(chan int, 50)
	go func(c <-chan int) {

		fmt.Println(<-chm)

	}(chm)
	go func(c chan<- int) {
		chm <- 41

	}(chm)
}
