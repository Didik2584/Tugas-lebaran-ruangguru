package main

import "fmt"

func countElements(slice []interface{}) int {
    return len(slice)
}

func main() {
    input1 := []interface{}{1, 2, 3, 4, 5}
    input2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    input3 := []interface{}{1, 1, 1, 5, 5, 5}
    input4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    input5 := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
    input6 := []interface{}{true, false, true, false, true, false}

    fmt.Println(countElements(input1))
    fmt.Println(countElements(input2))
    fmt.Println(countElements(input3))
    fmt.Println(countElements(input4))
    fmt.Println(countElements(input5))
    fmt.Println(countElements(input6))
}
