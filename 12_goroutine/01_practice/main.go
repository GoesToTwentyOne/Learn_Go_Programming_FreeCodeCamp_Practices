package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// control go routines with wait groups

	wg.Add(2)

	go gotest()

	// same go routines anonymous func
	go func() {
		fmt.Println("from anonymous function")
		wg.Done()

	}()
	wg.Wait()
	fmt.Println("all exist")
	// race condition
	raceConditionCreate()
	racePrevent()

}
func gotest() {
	fmt.Println("Hello form go test")
	wg.Done()
}

func raceConditionCreate() {
	var wgr sync.WaitGroup
	incrementor := 0

	gr := 100
	wgr.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := incrementor
			v++
			runtime.Gosched()
			incrementor = v
			fmt.Println("enter", incrementor)
			wgr.Done()

		}()

	}
	wgr.Wait()
	fmt.Println("exit", incrementor)
	fmt.Println("all exist")

}
func racePrevent() {
	var wgr sync.WaitGroup
	incrementor := 0

	gr := 100

	wgr.Add(gr)
	var mu sync.Mutex
	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			v++

			incrementor = v
			fmt.Println("enter", incrementor)
			wgr.Done()
			mu.Unlock()

		}()

	}
	wgr.Wait()
	fmt.Println("exit", incrementor)
	fmt.Println("all exist")

}
