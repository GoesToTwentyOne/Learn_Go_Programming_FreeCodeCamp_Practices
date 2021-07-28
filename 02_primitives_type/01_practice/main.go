package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------------integer--------------------------------------")
	//zero values
	var i int
	fmt.Printf("%T %v \n", i, i)
	//operations
	ia := 10
	ib := 20
	//arithmatic operation
	fmt.Println(ia + ib)
	fmt.Println(ia - ib)
	fmt.Println(ia * ib)
	fmt.Println(ia / ib)
	fmt.Println(ia % ib)
	//bitwise operation
	ba := 10
	bb := 2
	fmt.Println(ba & bb)
	fmt.Println(ba | bb)
	fmt.Println(ba ^ bb)
	fmt.Println(ba &^ bb)
	//bit shifting operation
	a := 1
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

	fmt.Println("-------------------------------------------float--------------------------------------")
	//zero values
	var f float32
	fmt.Printf("%T %v \n", f, f)
	fa := 3e3
	fb := 3.79e3
	//arithmatic operation
	fmt.Println(fa + fb)
	fmt.Println(fa - fb)
	fmt.Println(fa * fb)
	fmt.Println(fa / fb)
	fmt.Println("-------------------------------------------complex--------------------------------------")
	//zero values
	var c complex64
	fmt.Printf("%T %v \n", c, c)
	ca := 2 + 3i
	cb := complex(3, 4)
	//arithmatic operation
	fmt.Println(ca + cb)
	fmt.Println(ca - cb)
	fmt.Println(ca * cb)
	fmt.Println(ca / cb)
	//built it func
	fmt.Println(real(ca))
	fmt.Println(imag(ca))
	fmt.Println("-------------------------------------------string--------------------------------------")
	//zero values
	var s string
	fmt.Printf("%T %v \n", s, s)
	sa := "I'am Nihad Hossain"
	fmt.Println(sa)
	//immutable
	//s[2]="HH"
	//collection of byte slice
	fmt.Println([]byte(sa))
	fmt.Println("-------------------------------------------rune--------------------------------------")
	//zero values
	var r rune
	fmt.Printf("%T %v \n", r, r) //alias of int32
	ra := 'N'
	fmt.Println(ra)

}
