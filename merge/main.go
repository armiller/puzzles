package main

import (
	"fmt"
)

func main() {
	a := []int{3, 5, 7}
	b := []int{3, 3, 4, 6, 10, 10}
	c := merge(a, b)
	fmt.Println(c)
}

func merge(a, b []int) []int {
	var combined []int
	x, y := 0, 0
	for x < len(a) && y < len(b) {
		if a[x] < b[y] {
			combined = append(combined, a[x])
			x++
		} else {
			combined = append(combined, b[y])
			y++
		}
	}
	combined = append(combined, a[x:]...)
	combined = append(combined, b[y:]...)

	return combined
}
