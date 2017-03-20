package main

import (
	"fmt"
)

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	var buf string
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			buf += "Fizz"
		}
		if i%5 == 0 {
			buf += "Buzz"
		}
		if len(buf) == 0 {
			fmt.Printf("%d", i)
		}
		fmt.Println(buf)
		buf = ""
	}
}
