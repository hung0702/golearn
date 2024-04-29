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
	fmt.Println(a, bInt, c, d, e, f, 5+e%int(c), g && h || h)
	fmt.Print(&a) // print pointer
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
