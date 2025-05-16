package main

import "fmt"

type Order struct {
	ID     int
	Amount float32
	Status string
}

func CancleOrder(order *Order) {
	order.Status = "cancelled"
}

func main() {
	order := Order{
		ID:     1,
		Amount: 1200.344,
		Status: "pending",
	}
	fmt.Println("Before : ", order)
	CancleOrder(&order)
	fmt.Println("After : ", order)
}
