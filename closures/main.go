package main

import "fmt"

func activeGiftCard() func() int {
	count := 10
	return func() int {
		count--
		return count
	}
}

func main() {
	countVlue := activeGiftCard()
	fmt.Println(countVlue())
	fmt.Println(countVlue())
}
