package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}
type receiver struct{}

func (r receiver) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err

}

//method sets
type person struct {
	name string
	age  int
}

func (p person) speak() string {
	return fmt.Sprintf("name %s age %d", p.name, p.age)

}
func (p *person) speak2() string {
	return fmt.Sprintf("name %s age %d", p.name, p.age)

}

type human interface {
	speak() string
}
type human2 interface {
	speak2() string
}

func saying1(h human) {
	fmt.Println(h.speak())

}
func saying2(h human2) {
	fmt.Println(h.speak2())

}

func main() {
	wri := receiver{}
	wri.Write([]byte("Hello"))
	//type switch
	var x interface{} = 21
	switch x.(type) {
	case int:
		fmt.Println("integer")
	case []int:
		fmt.Println("slice of int")
	case bool:
		fmt.Println("bool")
	case float32:
		fmt.Println("float")
	default:
		fmt.Println("other types")
	}

	//method sets
	per := person{
		name: "Nihad",
		age:  21,
	}
	//normal recever (T *T)
	saying1(per)
	saying1(&per)
	//pointer recever (*T)
	//saying2(per)  don't work
	saying2(&per)

}
