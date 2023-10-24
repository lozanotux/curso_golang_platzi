package main

import "fmt"

func main() {
	// Maps
	m := make(map[string]int)

	m["Jhon"] = 14
	m["Doe"] = 21

	fmt.Println(m)

	// Walk through map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Find a value
	find, is_ok := m["Unknown"]
	fmt.Println(find, is_ok)
}
