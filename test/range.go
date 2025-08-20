package main

import "fmt"

func main() {
	var sliceWord string = "abc"
	fmt.Printf("%T %v\n", sliceWord, sliceWord)
	for _, char := range sliceWord {
		fmt.Printf("%T %c\n", char, char)
	}
}
