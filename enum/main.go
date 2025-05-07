package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func (os OrderStatus) String() string {
	switch os {
	case Received:
		return "Recived"
	case Confirmed:
		return "Confirmed"
	case Prepared:
		return "Prepared"
	case Delivered:
		return "Delivere"
	default:
		return fmt.Sprintf("Unknow OderStatus: %d", os)
	}
}

func changeOrderStatus(status OrderStatus) {
	switch status {
	case Received:
		fmt.Println("Order is now Received .")
	case Confirmed:
		fmt.Println("Order is now Confirmed .")
	case Prepared:
		fmt.Println("Order is now Prepared .")
	case Delivered:
		fmt.Println("Oder is now Delivered .")
	default:
		fmt.Println("Unknown Order Status .")
	}
}
func main() {
	changeOrderStatus(Prepared)
}
