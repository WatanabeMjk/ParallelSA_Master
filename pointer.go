package main

import "fmt"

func main() {
	var testInt = 0
	test(&testInt)
	fmt.Printf("test : %d", testInt)
}

func test(pointerInt *int) {
	*pointerInt = 100
}
