package main

import "fmt"

//declare const typed constant
const sha int32 = 21

//declare const typed constant
const c2 = 21

const (
	_ = iota
	b
	c
)

func main() {
	//allow shadowing and change the similar type
	const sha int64 = 73
	fmt.Printf("%T %v\n", sha, sha)
	//allow interoprate same type
	var x int64
	fmt.Println(sha + x)

	fmt.Printf("%T %v\n", c2, c2)

	//allow interoprate similar type type and type changed at the compile time
	var y int64 = 5
	fmt.Println(c2 + y)
	fmt.Printf("%v %T \n", c2+y, c2+y)

	//Enumerated constant and expression
	fmt.Println(b, c)

}
