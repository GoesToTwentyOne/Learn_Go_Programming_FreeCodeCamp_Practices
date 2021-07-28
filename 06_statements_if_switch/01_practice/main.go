package main

import "fmt"

func main() {
	//if else-if
	x := 10
	if x == 2 {
		fmt.Println("two")
	} else if x == 3 {
		fmt.Println("three")
	} else if x == 10 {
		fmt.Println("ten")
	} else {
		fmt.Println("others")
	}

	// simple switch
	switch x := 10; x {
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 10:
		fmt.Println("ten")
		fallthrough

	case 5:
		fmt.Println("five")
	default:
		fmt.Println("others")
	}

	// simple type
	var xm interface{} = 21

	switch xm.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("int")
	case bool:
		fmt.Println("int")
	default:
		fmt.Println("others")

	}

}
