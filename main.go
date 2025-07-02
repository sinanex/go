package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(add(18, 34))

	var age int = 19
	fmt.Println(age)

	var name string = "Sinan"
	fmt.Println(message(name))

	nameprint() // ✅ now this will work
}

func add(a int, b int) string {
	sum := a + b
	return fmt.Sprintf("Result: %d + %d = %d", a, b, sum)
}

func message(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func nameprint() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Name is %s\n", name)
}
