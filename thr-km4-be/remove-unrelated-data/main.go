package main

import "fmt"

func removeUnrelatedData(mapData map[string]interface{}, key string) map[string]interface{} {
	delete(mapData, key)
	return mapData
}

func main() {
	map1 := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	result1 := removeUnrelatedData(map1, "address")
	fmt.Println(result1)

	map2 := map[string]interface{}{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	result2 := removeUnrelatedData(map2, "jump")
	fmt.Println(result2)

	map3 := map[string]interface{}{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	result3 := removeUnrelatedData(map3, "tiga")
	fmt.Println(result3)
}
