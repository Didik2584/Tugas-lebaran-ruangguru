package main

import "fmt"

func removeElement(array []interface{}, position string) []interface{} {
    if position == "left" {
        return array[1:]
    } else {
        return array[:len(array)-1]
    }
}

func main() {
    input1 := []interface{}{1, 2, 3, 4, 5}
    output1 := removeElement(input1, "left")
    fmt.Println(output1) // [2 3 4 5]

    input2 := []interface{}{1, 2, 3, 4, 5}
    output2 := removeElement(input2, "right")
    fmt.Println(output2) // [1 2 3 4]

    input3 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    output3 := removeElement(input3, "left")
    fmt.Println(output3) // [Budi Joko Tono]

    input4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    output4 := removeElement(input4, "right")
    fmt.Println(output4) // [Edo Budi Joko]
}
