package main

import "fmt"

func main() {
	var word string = "abc"
	sliceWord := []rune(word)
	fmt.Printf("%T %v\n", sliceWord, sliceWord)
	for _, char := range sliceWord {
		fmt.Printf("%T %c\n", char, char)
	}
}
