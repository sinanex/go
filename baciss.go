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

	add(12,45)
}

func add(a int, b int) string {
	sum := a + b
	return fmt.Sprintf("result %d + %d  = %d",a,b, sum)
}

func message(name string) string {
	return fmt.Sprintf("hello %s", name)
}
