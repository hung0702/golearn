package main

import (
	"log"
	"os"
)

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

/*
Slicing arrays: slices reference an array, and are resizable. Can make, copy,
append and slice again.
*/

// func main() {
// 	OddNums := [4]int{1, 3, 5, 7} // shorthand init/decl array elements
// 	// for _, value := range OddNums { // no iterator; range returns value only
// 	// 	fmt.Println(value)
// 	// }
// 	// fmt.Println(OddNums[2])
// 	for i, value := range OddNums { // with iterator; range returns value + it
// 		fmt.Println(i, value)
// 	}
// 	arrayToSlice := []int{5, 4, 3, 2, 1} // array; no fixed length
// 	sliced := arrayToSlice[3:5] // slice of an array; 2nd index not inclusive
// 	fmt.Println(sliced)
// 	slice2 := make([]int, 5, 10) // make a slice (type, length, capacity)
// 	copy(slice2, arrayToSlice) // copy a slice (destination, source)
// 	fmt.Println(arrayToSlice)
// 	slice3 := append(slice2, 6, 7, 8) // append to a slice
// 	fmt.Println(slice3)
// }

// func main() {

// 	// StudentAge := make(map[string]int)

// 	// StudentAge["Helen"] = 30
// 	// StudentAge["Hung"] = 33
// 	// StudentAge["Hieu"] = 35
// 	// StudentAge["Hoa"] = 37
// 	// fmt.Println(StudentAge)          // prints entire map
// 	// fmt.Println(StudentAge["Helen"]) // prints key's value

// 	kitties := map[string]map[string]string{
// 		"Cali": map[string]string{
// 			"coat":   "Calico",
// 			"weight": "300g",
// 		},
// 		"White Paws": map[string]string{
// 			"coat":   "Black and White",
// 			"weight": "338g",
// 		},
// 		"All Black": map[string]string{
// 			"coat":   "Black",
// 			"weight": "302g",
// 		},
// 		"Tabby": map[string]string{
// 			"coat":   "Tabby",
// 			"weight": "342g",
// 		},
// 	}

// 	if temp, kitten := kitties["Cali"]; kitten {

// 		fmt.Println("Cali", temp["coat"], temp["weight"])
// 	}
// }

/*
Functions
*/

// func main() {

// 	x, y := 50, 600
// 	fmt.Println(add(x, y))

// }

// func add(x, y int) int { // name, parameters (type, can be multi), return type
// 	return x + y
// }

/*
Recursion via factorial (function calls itself for next member and has exit condition)
*/

// func main() {
// 	fmt.Println(factorial(5))
// }

// func factorial(x int) int {
// 	if x == 0 {
// 		return 1
// 	}
// 	return x * factorial(x-1)
// }

/*
Defer: defers fxn exec until surrounding function returns. LIFO. Generally:
cleans up resources e.g. db conxns.
Recover: return to normal after panic; stops panic, returns value that caused
panic. Generally: used with defer in goroutines.
Panic: Like an exception; stops fxn, starts panic; deferred fxns still run.
*/

// func main() {

// 	defer FirstRun()
// 	SecondRun()

// }

// func FirstRun() {
// 	fmt.Println("deferred First")
// }

// func SecondRun() {
// 	fmt.Println("Second")
// }

// func main() {
// 	// fmt.Println(div(3, 9))
// 	fmt.Println(div(3, 0))
// }

// func div(x, y int) float64 {
// 	defer func() {
// 		if y == 0 {
// 			fmt.Println("Division by zero")
// 		}
// 		if recover() != nil { // conditionally print recover()
// 			fmt.Println(recover())
// 		}
// 	}()
// 	solution := float64(x) / float64(y)
// 	return solution
// }

// func main() {
// 	fmt.Println(addemup(1, 2, 3, 4, 5))
// }

// func addemup(args ...int) int { // variadic fxns take any number of arguments

// 	total := 0
// 	for _, value := range args { // range returns index and value
// 		total += value
// 	}
// 	return total
// }

/*
Structures: arrays but allows different data types; can define objects
Interfaces: collection of methods; can be used to define behavior.
*/

/*
The animal interface allows for kitten and dog structs, which both have
different variables and different weight units. For each struct, a specific
metrictooz function will convert the weight based on known input unit. The
printWeightInOz function only serves to consistently print the weight, easing
future maintenance.

Use PascalCase for public scope, camelCase for private scope.

An enhanced implementation of this below function would use embedded struct for
animal since name is repeated, and only specify unique variables for each struct
(animal).
*/

// func main() {
// 	cali := Kitten{Name: "Cali", Coat: "Calico", grams: 300}
// 	imaginary := Dog{Name: "Spot", kilograms: 5}

// 	fmt.Println(cali)
// 	printWeightInOz(cali)

// 	fmt.Println(imaginary)
// 	printWeightInOz(imaginary)
// }

// type Animal interface {
// 	metrictooz() float32
// }

// type Kitten struct {
// 	Name  string
// 	Coat  string
// 	grams int
// }

// type Dog struct {
// 	Name      string
// 	kilograms int
// }

// func (k Kitten) metrictooz() float32 {
// 	return float32(k.grams) * 0.03527396
// }

// func (d Dog) metrictooz() float32 {
// 	return float32(d.kilograms) * 35.27396
// }

// func printWeightInOz(a Animal) {
// 	fmt.Println("Weight in oz: ", a.metrictooz())
// }

/*
File handling: read, write, create, delete, rename, etc.
*/

func main() {

	file, err := os.Create("example.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi, I'm a file and I'm here to stay!")
	file.Close()

}
