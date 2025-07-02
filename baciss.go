package main

import "fmt"

func main() {

	fmt.Println(add(18, 34))
	//data types

	var age int = 19

	//uint -only positive nummbers

	fmt.Println(age)

	var name string = "sinan"
	fmt.Println(message(name))
}

func add(a int, b int) int {
	return a + b
}

func message(name string) string {
	return fmt.Sprintf("hello %s", name)
}
