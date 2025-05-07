package main

import (
	"fmt"
	"time"
)

type customer struct {
	name   string
	mobile string
}
type order struct {
	id        string
	ammount   float32
	status    string
	createdAt time.Time
	customer
}

func newOrder(id string, ammount float32, status string, cust customer) *order {
	myOrder := order{
		id:        id,
		ammount:   ammount,
		status:    status,
		createdAt: time.Now(),
		customer:  cust,
	}
	return &myOrder
}

func (o *order) updateAmount(newAmout float32) {
	o.ammount = newAmout
}

// recomanded
func (o *order) updateNameOne(newName string) *order {
	o.customer.name = newName
	return o
}

// not recomaned
func (o *order) updateNameTwo(newName string) {
	o.customer.name = newName
}
func main() {
	newcustomer := customer{
		name:   "Noman",
		mobile: "01883217001",
	}
	myoder := newOrder("order-1", 12.33, "pending", newcustomer)
	// myoder.customer = newcustomer
	fmt.Println(myoder)
	myoder.updateAmount(1234.45)
	fmt.Println(myoder)

}
