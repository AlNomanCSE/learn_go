package main

import (
	"fmt"
	"time"
)

type Printable interface {
	PrintDetails() string
}
type customer struct {
	name   string
	mobile string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (c customer) PrintDetails() string {
	return fmt.Sprintf("Customer:%s , Mobile :%s ", c.name, c.mobile)
}
func (o order) PrintDetails() string {
	return fmt.Sprintf("Order ID: %s, Amount: %.2f, Status: %s", o.id, o.amount, o.status)
}
func printEntity(p Printable) {
	fmt.Println(p.PrintDetails())
}

// func main() {
// 	cust := customer{name: "Noman", mobile: "01883217001"}
// 	printEntity(cust)
// }
