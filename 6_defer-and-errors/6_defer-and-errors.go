package main

import (
	"fmt"
)

/*
Defer: defers fxn exec until surrounding function returns.
	* LIFO
	* Generally: cleans up resources e.g. db conxns.

Recover: return to normal after panic
	* stops panic, returns value that caused it
	* Generally: used with defer in goroutines.

Panic: Like an exception
	* stops fxn
	* starts panic
	* deferred fxns still run
*/

func main() {

	defer FirstRun()
	SecondRun()

	fmt.Println("div/0 handling")
	fmt.Println(div(3, 9))
	fmt.Println(div(3, 0))
}

func FirstRun() {
	fmt.Println("deferred First")
}

func SecondRun() {
	fmt.Println("Second")
}

func div(x, y int) float64 {
	defer func() {
		if y == 0 {
			fmt.Println("Division by zero")
		}
		if recover() != nil { // conditionally print recover()
			fmt.Println(recover())
		}
	}()
	solution := float64(x) / float64(y)
	return solution
}
