package main

import "fmt"

func multiplier(factor int) func(int) int {

	return func(i int) int {
		return i * factor
	}
}

func main() {

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(2), triple(3))

}
