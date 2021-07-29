package main

import "fmt"

func main() {
	// creating pointers
	//dereferencing
	var a int
	var b *int
	fmt.Printf("%T  \t %v \t %T \t %v\n", a, a, b, b)
	b = &a
	*b = 21
	fmt.Printf("%T  \t %v \t %T \t %v\n", a, a, b, *b)

	// creating pointers like object
	type person struct {
		name string
		age  int
	}
	//process 1
	// var p1 *person
	// p1 = new(person)
	// (*p1).name = "Setaelo"
	// (*p1).age = 21
	// fmt.Println(p1)

	p2 := &person{name: "Nihad", age: 21}
	fmt.Println(p2)
	p3 := new(person)
	fmt.Println(p3)

	//types with internal pointers
	// pointed to underlying data types
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(xi)
	xi2 := xi
	xi2[0] = 21
	fmt.Println(xi, xi2)
	// don't pointed to underlying data types
	xia := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(xia)
	xia2 := xia
	xia2[0] = 21
	fmt.Println(xia, xia2)
	//types with internal pointers
	// pointed to underlying data types
	m := map[string]int{
		"Math":    21,
		"English": 11,
		"ICT":     10,
	}
	fmt.Println(m)
	m1 := m

	m1["ICT"] = 200
	fmt.Println(m)
	fmt.Println(m1)

}
