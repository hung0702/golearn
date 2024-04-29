package main

import "fmt"

/*
For loops to iterate. Can nest.
If/Switch statements for conditional logic
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

	fmt.Print("\n")
	fmt.Println("if statement")
	ifStatement()

	fmt.Print("\n")
	fmt.Println("switch statement")
	switchStatement()

	fmt.Print("\n")
	fmt.Println("switch statement with ranges")
	switchStatementWithRanges()
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
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func ifStatement() {

	age := 17

	if age <= 17 {
		fmt.Println("You no vote!")
	} else if age >= 18 {
		fmt.Println("You vote now!")
	} else {
		fmt.Println("How old is you?")
	}
}

func switchStatement() {
	// shorthand for If, Else If, Else statements; only exact match
	age := 18

	switch age {
	case 17:
		fmt.Println("You no vote!")
	case 18:
		fmt.Println("You vote now!")
	default:
		fmt.Println("How old is you?")
	}
}

func switchStatementWithRanges() {
	// shorthand for  statements using ranges
	age := 18
	switch {
	case age <= 17:
		fmt.Println("You no vote!")
	case age >= 18:
		fmt.Println("You vote now!")
	default:
		fmt.Println("How old is you?")
	}
}
