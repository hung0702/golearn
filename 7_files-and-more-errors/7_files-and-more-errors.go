package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Sum of 1 to 5 =", addemup(1, 2, 3, 4, 5))

	createFileExample()
}

func addemup(args ...int) int { // variadic fxns take any number of arguments

	total := 0
	for _, value := range args { // range returns index and value
		total += value
	}
	return total
}

/*
File handling and error logging
*/

func createFileExample() {

	file, err := os.Create("example.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi, I'm a file and I'm here to stay!")
	file.Close()

	stream, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	s1 := string(stream)

	fmt.Println(s1)
}
