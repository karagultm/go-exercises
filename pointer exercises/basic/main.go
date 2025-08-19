// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var pointer *computer

	// compare the pointer variable to a nil value, and say it's nil
	if pointer == nil {
		fmt.Print("Pointer is nil\n")
	}

	// create an apple computer by putting its address to a pointer variable
	pointer = &computer{"apple"}

	// put the apple into a new pointer variable
	newPointer := pointer

	// compare the apples: if they are equal say so and print their addresses
	if pointer.brand == newPointer.brand {
		fmt.Printf("old pointer adress %p new pointer adress %p\n", pointer, newPointer)
	}

	// create a sony computer by putting its address to a new pointer variable
	newPointer = &computer{"sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if pointer.brand != newPointer.brand {
		fmt.Printf("apple pointer adress %p sony pointer adress %p\n", pointer, newPointer)

	}

	// put apple's value into a new ordinary variable
	newApple := *pointer

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple pointer variable adress %p points to %p new varibale adress %p\n", &pointer, pointer, &newApple)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *pointer == newApple {
		fmt.Printf("tehy are equal\n")
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("values apple %s new variable %s\n", *pointer, newApple)

	// create a new function: change
	// the func can change the given computer's brand to another brand
	change(newPointer, "hp")
	fmt.Printf("new hp : %s\n", newPointer.brand)

	// change sony's brand to hp using the func — print sony's brand

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses

	// print the returned value
	fmt.Printf("appleVal                  : %+v\n", valueOf(pointer))
	fmt.Printf("appleVal                  : %+v\n", valueOf2(*pointer))

	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}
func change(v *computer, s string) {
	v.brand = s
}

func valueOf2(v computer) (s string) {
	s = v.brand
	return
}

// write a func that returns the value that is pointed by the given *computer
func valueOf(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer
func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
