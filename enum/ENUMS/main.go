package main

import "fmt"

type PaymentMethod int

const (
	CreditCard PaymentMethod = iota
	DebitCard
	BankTransfer
	DigitalWallet
	Cash
)

func (pm PaymentMethod) String() string {
	switch pm {
	case CreditCard:
		return "Credit Card"
	case DebitCard:
		return "Debit Card"
	case BankTransfer:
		return "Bank Transfer"
	case DigitalWallet:
		return "Digital Wallet"
	case Cash:
		return "Cash"
	default:
		return fmt.Sprintf("Unknown Payment Method: %d", pm)
	}
}
func processPayment(method PaymentMethod, amount float64) {
	fmt.Printf("Processing payment of $%.2f via %s\n", amount, method)
}
func main() {
	processPayment(CreditCard, 42.50)
}
