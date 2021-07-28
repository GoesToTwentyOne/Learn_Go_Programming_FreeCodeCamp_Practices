package main

import "fmt"

func main() {
	//map
	m := map[string]string{
		"Dhaka":      "10M",
		"Chittagong": "5M",
		"Rangpur":    "2M",
	}
	fmt.Println(m)
	//invoke one field value
	fmt.Println(m["Dhaka"])
	//comma ok idiom
	_, ok := m["Dha"]
	if ok {
		delete(m, "Dhaka")
		fmt.Println("deleted successfully")

	} else {
		fmt.Println("value doesn't exist")
	}

	//embedded type struct
	type person struct {
		name string
		age  int
	}
	type police struct {
		person
		gun string
	}
	poli := police{
		person: person{
			name: "Anrghoo",
			age:  21,
		},
		gun: "Yes",
	}
	fmt.Println(poli)

	// anonymous struct
	po := struct {
		name string
		age  int
	}{
		name: "Gtriko",
		age:  21,
	}
	fmt.Println(po)

}
