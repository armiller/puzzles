package main

import (
	"fmt"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func rotate(s string, rot int) string {
	rot %= 26
	b := []byte(s)
	for i, c := range b {
		b[i] = alphabet[(int(c-'a')+rot)%26]
	}
	return string(b)
}

func main() {
	fmt.Println(rotate("ace", 2))
}
