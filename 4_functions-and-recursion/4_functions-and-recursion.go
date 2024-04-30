package main

import "fmt"

/* specific syntax*/

func main() {

	x, y := 50, 600

	fmt.Println(x, "+", y, "=", add(x, y))

	fmt.Println("\n5! =", factorial(5))

}

func add(x, y int) int { // name, parameters (type, can be multi), return type
	return x + y
}

/*
Recursion via factorial (function calls itself for next member and has exit condition)
*/

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
