package main

import "fmt"

func main() {
	//for clause
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//for clause two
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	// for single condition with label
label:
	for i := 0; i < 4; i++ {
		fmt.Println("Outer index ", i)
		for j := 0; j < 4; j++ {
			fmt.Println("\t", i*j)
			if i*j >= 3 {
				break label
			}
		}
	}
	// for single condition 2
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//for infinity
	x := 0
	for {
		x++
		fmt.Println(x)
		if x == 5 {
			break
		}
	}
	//range clause
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
