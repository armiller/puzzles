package main

import "fmt"
import "strconv"

func main() {
	test := "aaabbbccbddd"
	result := compress(test)
	fmt.Println(result)
}

func compress(uncompressed string) string {
	prev := 0
	var letter string
	var next string
	var compressed string
	for index := 1; index < len(uncompressed); index++ {
		letter = string(uncompressed[index-1])
		next = string(uncompressed[index])
		if letter != next {
			count := index - prev
			compressed += string(letter)
			compressed += strconv.Itoa(count)
			prev = index
		}
	}
	compressed += string(letter)
	compressed += strconv.Itoa(len(uncompressed) - prev)
	return compressed
}
