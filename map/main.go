package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap1 := map[string]int{
		"price": 12300,
		"name":  111,
	}
	myMap2 := make(map[string]int)
	fmt.Println(myMap1, myMap2)
	maps.Copy(myMap2, myMap1)
	fmt.Println(myMap1, myMap2)
}
