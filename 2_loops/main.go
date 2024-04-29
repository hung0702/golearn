package main

import "fmt"

/*
 For loops to iterate. Can nest.
*/

func main() {
	fmt.Println("for loop")
	forLoop()

	fmt.Print("\n")
	fmt.Println("while loop")
	whileLoop()

	fmt.Print("\n")
	fmt.Println("nested while loop")
	nestedWhileLoop()
}

func forLoop() {
	// init/decl iterator; limit; incrementation
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
}

func whileLoop() {
	i := 1      // init/decl iterator
	for i < 5 { // limit
		fmt.Println(i)
		i++ // incrementation
	}
}

func nestedWhileLoop() {
	for i := 1; i < 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
