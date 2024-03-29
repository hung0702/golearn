package main

import "fmt"

// "strconv"

/*
Must Initialize a variable; can declare simultaneously and re-declare in different scope. Var types are immutable, and inferred. Constants are declared with const. Shorthand for declaring and initializing multiple vars.
*/

// func main() {
// 	var a = 1     // type inference; once declared, immutable
// 	const b = "2" // contant; strings cast to int with strconv.Atoi()
// 	bInt, err := strconv.Atoi(b) // strings don't easily convert,
// 	if err != nil {              // function to cast returns int and err
// 		fmt.Println("not an int", err)
// 		return
// 	}

// 	var (
// 		c float32 = 3.0000001 // loses precision past 6 decimal places
// 		d         = 4
// 	)
// 	e, f := 5, 6 // shorthand for declaring and initializing multiple vars
// 	g := true
// 	h := false
// 	fmt.Println(a, bInt, c, d, e, f, 5+e%int(c), g && h || h)
// 	fmt.Print(&a)
// }

/*
Pointers are used to change the value of a variable in a function.
*/

// func main() {
// 	x := 10
// 	changeValue(&x)
// 	print(x)
// }

// func changeValue(x *int) {
// 	*x = 7
// }

/*
Printf is used to format output.
*/

// func main() {
// 	wbo := true
// 	const pi = 3.14
// 	fmt.Printf("%.2f \n", pi) // %f for float, .2 for precision
// 	fmt.Printf("%T \n", pi)   // print the var type
// 	fmt.Printf("%t \n", wbo)  // printing boolean values requires lowercase t
// 	fmt.Printf("%b \n", 25)   // binary
// 	fmt.Printf("%c \n", 34)   // characters of the code value
// 	fmt.Printf("%x \n", 34)   // hex code
// 	fmt.Printf("%e \n", pi)   // scientific notation
// }

/*
 For loops are used to iterate through a block of code. Nested for loops are used to iterate through a block of code within a block of code.
*/

// func main() {
// 	for i := 1; i < 5; i++ { // init/decl iterator; limit; incrementation
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	// while loop version
// 	i := 1      // init/decl iterator
// 	for i < 5 { // limit
// 		fmt.Println(i)
// 		i++ // incrementation
// 	}
// }

// func main() {
// 	// while loop version
// 	for i := 1; i < 5; i++ {
// 		for j := 1; j < i; j++ {
// 			fmt.Printf("*")
// 		}
// 		fmt.Println()
// 	}
// }

/*
If/Switch statements for conditional logic
*/

// func main() {
// 	age := 18
// 	if age >= 18 {
// 		fmt.Println("You vote!")
// 	} else {
// 		fmt.Println("You no vote!")
// 	}
// }

// func main() {
// 	age := 18

// 	switch age {
// 	case 17:
// 		fmt.Println("You no vote!")
// 	case 18:
// 		fmt.Println("You vote now!")
// 	default:
// 		fmt.Println("How old is you?")
// 	}
// }

/*
Arrays: all elements same type (not immut), fixed length, and zero-indexed.
*/

// func main() {
// 	var OddNums [4]int

// 	OddNums[0] = 1
// 	OddNums[1] = 3
// 	OddNums[2] = 5
// 	OddNums[3] = 7

// 	fmt.Println(OddNums[2])
// }

func main() {
	OddNums := [4]int{1, 3, 5, 7} // shorthand init/decl array elements

	// for _, value := range OddNums { // no iterator; range returns value only
	// 	fmt.Println(value)
	// }
	// fmt.Println(OddNums[2])

	for i, value := range OddNums { // with iterator; range returns value + it
		fmt.Println(i, value)
	}

	numSlice := []int{5, 4, 3, 2, 1} // slice; no fixed length

	sliced := numSlice[3:5] // slice of a slice; 2nd index not inclusive

	fmt.Println(sliced)

	slice2 := make([]int, 5, 10) // make a slice (type, length, capacity)

	copy(slice2, numSlice) // copy a slice (destination, source)

}
