package main

import "fmt"

func main() {
	name := "Bob"

	fmt.Println(len(name))
	fmt.Println(name + " is ok")

	const pi = 3.14159
	fmt.Printf("")
}

// "strconv"

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

// func main() {

// 	x := 10

// 	changeValue(&x)
// 	print(x)
// }

// func changeValue(x *int) {

// 	*x = 7
// }
