package main

import "fmt"

var (
	names     []string
	distances []int
	data      []byte
	ratios    []float64
	alives    []bool
)

func main() {
	names = []string{"nail", "eröl", "semih"}
	distances = []int{12, 321, 52, 211, 56}
	data = []byte{'I', 'N', 'A', 'N', 'C'}
	ratios = []float64{21.2, 31.8}
	alives = []bool{true, false, false, true}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Print("The length of the distances and the data slices are the same.")
	}
}
