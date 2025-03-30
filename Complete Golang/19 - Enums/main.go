package main

import "fmt"

//enumerated types
type OrderStatus string

const (
	Recieved  OrderStatus = "Received"
	confirmed             = "confirmed"
	Delivered             = "Delivered"
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ", status)
}

func main() {
	ChangeOrderStatus(confirmed)
}
