package main

import (
	"fmt"
	"strconv"
)

/*
* Variables and Constants: must be declared and initialized
	* Cannot initialize in different scope
* Var types are immutable, and inferred
* Two shorthand ways to declare and initialize a variable
* Constants have no shorthand
*/

func main() {
	var a = 1                    // type inference; immutable
	const b = "2"                // constants have no shorthand
	bInt, err := strconv.Atoi(b) // string conversion
	if err != nil {              // function to cast returns int and err
		fmt.Println("not an int", err)
		return
	}

	var ( // declare multiple vars more legibly
		c float32 = 3.0000001 // loses precision past 6 decimal places
		d         = 4
	)
	e, f := 5, 6 // shorthand to declare and initialize 1 or more vars
	g := true
	h := false

	fmt.Println(a, b, c, d, f) // println adds space after each, new line at the end
	fmt.Println(bInt)
	fmt.Println(5 + e%int(c))
	fmt.Println(g && h || h)

	fmt.Print(&a, "\n")   // print pointer
	fmt.Printf("%T\n", b) // print type of b
	PrintfExamples()
}

/*
Types
* string
* bool
* int
* int	int8	int16	int32	int64
	* rune - alias for int32
* uint	uint8	uint16	uint32	uint64	uintptr
	* byte - alias for uint8
* float32	float64
* complex64	complex128
*/

func PrintfExamples() { // Printf to format output
	// PascalCase exports vs camelCase
	wbo := true
	const pi = 3.14
	fmt.Printf("%.2f \n", pi) // %f for float, .2 for precision
	fmt.Printf("%T \n", pi)   // print the var type
	fmt.Printf("%t \n", wbo)  // printing boolean values requires lowercase t
	fmt.Printf("%b \n", 25)   // binary
	fmt.Printf("%c \n", 34)   // characters of the code value
	fmt.Printf("%x \n", 34)   // hex code
	fmt.Printf("%e \n", pi)   // scientific notation
}
