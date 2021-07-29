package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println(p.name, p.age)
	//value doesn't exist
	// p.name = "don't changing"
}
func (p *person) speak2() {
	fmt.Println(p.name, p.age)
	//value doesn't exist
	p.name = "changing value"
}
func main() {
	//parameters
	sum, sub := simpeSumSub(10, 10)
	fmt.Println("sum :", sum, "sub :", sub)
	//variadic sum
	fmt.Println(variadicSum(1, 2, 3, 34, 5, 6, 66, 6, 6))
	//pointer return
	fmt.Println(*ptrreturn(1, 2, 3, 4, 5, 6, 7, 9, 10)) // same line deref
	//anonymous func
	for i := 0; i < 10; i++ {
		func(x int) {
			fmt.Println(i)

		}(i)
	}
	// function as a types

	divide := func(x, y int) (int, error) {
		if x|y == 0 {
			return 0, fmt.Errorf("value %v %v is not exist the func", x, y)
		}
		return x / y, nil

	}
	fmt.Println(divide(5, 5))
	//method
	p := person{name: "Nihad", age: 21}
	p.speak()
	fmt.Println(p.name)
	p.speak2()
	fmt.Println(p.name)

}
func simpeSumSub(x, y int) (int, int) {
	return x + y, x - y

}

//variadic parameter
func variadicSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//pointer return
func ptrreturn(x ...int) *int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return &sum
}
