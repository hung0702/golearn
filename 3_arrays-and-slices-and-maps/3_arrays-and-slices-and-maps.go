package main

import "fmt"

func main() {
	fmt.Println("Array example: prints element 3")
	arrayExample()

	fmt.Println("\nSlice arrays with loop")
	sliceArraysWithLoop()

	fmt.Println("\nSlice arrays")
	sliceArrays()

	fmt.Println("\nMaps = dictionaries")
	mapExample()

	fmt.Println("\nNested maps")
	nestedMaps()
}

func arrayExample() {
	// all elements same type (not immut), fixed length, and zero-indexed
	var OddNums [4]int

	OddNums[0] = 1
	OddNums[1] = 3
	OddNums[2] = 5
	OddNums[3] = 7

	fmt.Println(OddNums[2])
}

func sliceArraysWithLoop() {
	// Slices reference an array, and are resizable.
	// Can make, copy, append, and re-slice.

	OddNums := [4]int{1, 3, 5, 7} // shorthand init/decl array elements
	// [arraySize]dataType{elements}
	fmt.Println(" value only")
	for _, value := range OddNums { // no identifier; range returns value only
		fmt.Println(value)
	} // blank identifier to ignore index and avoid unused variable i issue

	fmt.Println(" index + value")
	for i, value := range OddNums { // range now returns index + value
		fmt.Println(i, value)
	} // identifer to get index; must use i to avoid unused variable i issue
}

func sliceArrays() {
	fmt.Println("\nSlices include the first index given but not the last")
	arrayToSlice := []int{1, 3, 5, 7} // array w/o fixed length
	slice1 := arrayToSlice[1:3]       // slice; inclusive-exclusive
	fmt.Println(slice1)

	fmt.Println("\nCan make and copy slices")
	slice2 := make([]int, 4, 10) // make a slice (type, length, capacity)
	// length = intial size; capacity (optional); if omitted, capacity = length
	copy(slice2, arrayToSlice) // copy a slice (destination, source)
	fmt.Println(arrayToSlice)

	fmt.Println("\nCan append to slices")
	slice3 := append(slice2, 9, 11, 13) // append to a slice
	fmt.Println(slice3)
}

func mapExample() {

	example := make(map[string]int)

	example["A"] = 1
	example["B"] = 3
	example["C"] = 5
	example["D"] = 7
	fmt.Println(example)                     // prints entire map
	fmt.Println("A has value", example["A"]) // prints key's value
}

func nestedMaps() {
	kitties := map[string]map[string]string{
		"Cali": {
			"coat":   "Calico",
			"weight": "300g",
		},
		"White Paws": {
			"coat":   "Black and White",
			"weight": "338g",
		},
		"All Black": {
			"coat":   "Black",
			"weight": "302g",
		},
		"Tabby": {
			"coat":   "Tabby",
			"weight": "342g",
		},
	}

	fmt.Println(kitties)                                     // print entire map
	fmt.Println("Cali")                                      // print submap
	fmt.Println("Cali's weight:", kitties["Cali"]["weight"]) // prints value in submap
}
