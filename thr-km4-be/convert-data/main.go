package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	parts := strings.Split(data, ",")
	age, _ := strconv.Atoi(parts[1])
	return User{Name: parts[0], Age: age, Address: parts[2]}
}

func main() {
	user1 := ConvertData("Edo,20,Jakarta")
	fmt.Printf("{ name: \"%s\", age: %d, address: \"%s\" }\n", user1.Name, user1.Age, user1.Address) // Output: { name: "Edo", age: 20, address: "Jakarta" }

	user2 := ConvertData("Budi,30,Bandung")
	fmt.Printf("{ name: \"%s\", age: %d, address: \"%s\" }\n", user2.Name, user2.Age, user2.Address) // Output: { name: "Budi", age: 30, address: "Bandung" }
}
