package main

import "fmt"

func findMax(numbers []int) int {
	max := 0
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max
}
func loop() {
	numbers := []int{10, 35, 2, 99, 7}
	max := findMax(numbers)
	fmt.Println(max)
}
