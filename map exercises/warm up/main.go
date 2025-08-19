// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	one := map[string]string{}

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	var second map[int]bool

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	var third map[string][]string

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	var fourth map[int]map[int]int

	fmt.Printf("phones     : %#v\n", one)
	fmt.Printf("products   : %#v\n", second)
	fmt.Printf("multiPhones: %#v\n", third)
	fmt.Printf("basket     : %#v\n", fourth)
}
