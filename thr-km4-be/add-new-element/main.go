package main

import "fmt"

// Definisikan fungsi AddNewElement dan main seperti sebelumnya

func AddNewElement(array []interface{}, element interface{}, position string) []interface{} {
	switch position {
	case "up":
		return append([]interface{}{element}, array...)
	case "down":
		return append(array, element)
	default:
		return array
	}
}

func main() {
	array := []interface{}{1, 2, 3, 4, 5}
	element := 6

	result := AddNewElement(array, element, "up")
	fmt.Println(result) // [6 1 2 3 4 5]

	result = AddNewElement(array, element, "down")
	fmt.Println(result) // [1 2 3 4 5 6]
}
