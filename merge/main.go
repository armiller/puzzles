package main

import (
	"fmt"
)

func main() {
	a := []int{3, 5, 7}
	b := []int{3, 3, 4}
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
	if len(a)-x != 0 {
		combined = append(combined, a[x:]...)
	}
	if len(b)-y != 0 {
		combined = append(combined, b[y:]...)
	}
	return combined
}
