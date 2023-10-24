package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	text = strings.ToLower(text)
	var reversedText string

	for i := len(text) - 1; i >= 0; i-- {
		reversedText += string(text[i])
	}

	if text == reversedText {
		fmt.Println("Is palindrome")
	} else {
		fmt.Println("Is not palindrome")
	}
}

func main() {
	// Array (len & cap)
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice (and methods: slicing)
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 8)
	fmt.Println(slice)
	slice = append(slice, 9, 10)

	// Append new list
	newSlice := []int{11, 12, 13}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Walk through Slices
	message_slice := []string{"hola", "que", "hace"}

	/*
	 * To skip i variable, just add _ (underscore)
	 * for example:  for _, value := range slice {}
	 */
	for i, value := range message_slice {
		fmt.Println(i, value)
	}

	// Palindrome texts
	isPalindrome("ama")
	isPalindrome("Ama")
	isPalindrome("amor a roma")
	isPalindrome("texto aleatorio")
}
