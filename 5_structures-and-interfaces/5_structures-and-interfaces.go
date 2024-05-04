package main

import "fmt"

/*
Structures: arrays but allows different data types; can define objects
Interfaces: collection of methods; can be used to define behavior.
	* Reference type, so it's always pointing and never copied
*/

/*
The animal interface allows for kitten and dog structs, which both have
different variables and different weight units. For each struct, a specific
metrictooz function will convert the weight based on known input unit. The
printWeightInOz function only serves to consistently print the weight, easing
future maintenance.

Use PascalCase to export stuff, camelCase if that stuff should only be available in scope.

The embedded struct for animal for name allows reuse of shared methods

Weight is not stored the same for each struct, hence the conversion function
*/

func main() {
	cali := Kitten{Name: "Cali", Coat: "Calico", grams: 300}
	imaginary := Dog{Name: "Spot", kilograms: 5}

	fmt.Println(cali)
	printWeightInOz(cali)

	fmt.Println(imaginary)
	printWeightInOz(imaginary)
}

type Animal interface {
	metrictooz() float32
}
type Mammal struct {
	Name string
}

type Kitten struct {
	Name  string
	Coat  string
	grams int
}

type Dog struct {
	Name      string
	kilograms int
}

func (k Kitten) metrictooz() float32 {
	return float32(k.grams) * 0.03527396
}

func (d Dog) metrictooz() float32 {
	return float32(d.kilograms) * 35.27396
}

func printWeightInOz(a Animal) {
	fmt.Println("Weight in oz: ", a.metrictooz())
}
