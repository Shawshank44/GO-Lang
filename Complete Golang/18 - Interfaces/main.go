package main

import "fmt"

// PaymentProcessor interface
type PaymentProcessor interface {
	Pay(amount float64) string
}

// CreditCard struct implementing PaymentProcessor
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card (%s)", amount, cc.CardNumber)
}

// PayPal struct implementing PaymentProcessor
type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal (%s)", amount, pp.Email)
}

// ProcessPayment function using interface
func ProcessPayment(p PaymentProcessor, amount float64) {
	fmt.Println(p.Pay(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234-5678-9012-3456"}
	pp := PayPal{Email: "user@example.com"}

	ProcessPayment(cc, 100.50)
	ProcessPayment(pp, 200.75)
}
