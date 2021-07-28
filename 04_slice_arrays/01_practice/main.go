package main

import "fmt"

func main() {
	//array
	xa := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(xa)
	//array slicing
	fmt.Println(xa[1:])
	fmt.Println(xa[:3])
	//[immutable]
	//copies refer to different underlying data {don't reference}
	xaa := [7]int{1, 2, 3, 4, 5, 6, 7}
	// var xaa2 [7]int
	xaa2 := xaa
	xaa2[0] = 21
	fmt.Println(xaa2)
	fmt.Println(xaa)

	//slice
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(xi)
	//array slicing
	fmt.Println(xi[1:])
	fmt.Println(xi[:3])
	//[mutable]
	//copies refer to same underlying data { reference--pointed to the address}
	xii := []int{1, 2, 3, 4, 5, 6, 7}
	xii2 := xii
	xii2[1] = 50
	fmt.Println(xii2)
	fmt.Println(xii)

	//copies refer to same underlying data { reference--pointed to the address}
	xia := []int{1, 2, 3, 4, 5, 6, 7}
	xib := append(xia[1:3], xia[4:]...)
	fmt.Println(xia)
	fmt.Println(xib)
}
