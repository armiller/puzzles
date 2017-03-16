package main

import (
	"fmt"
)

func main() {
	a := []int{3}
	b := []int{2, 3, 4}
	c := merge(a, b)
	fmt.Println(c)
}

func merge(a, b []int) []int {
	var combined []int
	total := len(a) + len(b)
	x, y := 0, 0
	for i := 0; i < total; i++ {
		if x < len(a) && y < len(b) {
			if a[x] < b[y] {
				combined = append(combined, a[x])
				x++
			} else {
				combined = append(combined, b[y])
				y++
			}
		} else if x < len(a) {
			combined = append(combined, a[x])
			x++
		} else {
			combined = append(combined, b[y])
			y++
		}
	}
	return combined
}
