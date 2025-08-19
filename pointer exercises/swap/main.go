// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	a, b := 3.14, 6.28
	fmt.Printf("a: %f b: %f\n", a, b)
	swap(&a, &b)
	fmt.Printf("Swapped a: %f b: %f\n", a, b)

	c, d := &a, &b
	fmt.Printf("adress of c: %p d: %p\n", c, d)
	swapAdress(&c, &d)
	fmt.Printf("swapped adress of c: %p d: %p", c, d)

}

func swap(a, b *float64) {
	temp := *a
	*a = *b
	*b = temp
	// *a, *b = *b, *a
	//bu daha kısa yoluymuşş
}

func swapAdress(c, d **float64) {
	*c, *d = *d, *c
}
