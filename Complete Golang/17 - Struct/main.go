package main

import (
	"fmt"
	"time"
)

// // order structure
// type Order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time
// }

// // constructor method for stucture
// func newOrder(id string, amount float32, status string) *Order {
// 	myOrder := Order{
// 		id:        id,
// 		amount:    amount,
// 		status:    status,
// 		createdAt: time.Now(),
// 	}

// 	return &myOrder
// }

// func (o *Order) ChangeStatus(status string) { // method for struct
// 	o.status = status
// }

// struct embedding
type Customer struct {
	name  string
	phone string
}

type Delivery struct {
	vendor  string
	FOBtype string
}

type Orders struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

type Entry struct {
	Customer
	Delivery
	Orders
}

func (e *Entry) statusMGT(status string) {
	e.Orders.status = status
}

func main() {

	// order := Order{
	// 	id:     "1",
	// 	amount: 50.25,
	// 	status: "Received",
	// }
	// order.createdAt = time.Now()

	// fmt.Println(order.status)

	// fmt.Println("Order struct", order)

	// Myorder := newOrder("1", 500.25, "recieved")

	// fmt.Println(Myorder)

	// languages := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(languages)

	neworder := Entry{
		Customer: Customer{
			name:  "John",
			phone: "98765434210",
		},
		Delivery: Delivery{
			vendor:  "FEDEX",
			FOBtype: "Next-day-air",
		},
		Orders: Orders{
			id:        "AMX001",
			amount:    550.65,
			status:    "Received",
			createdAt: time.Now(),
		},
	}

	neworder.statusMGT("On the way")

	fmt.Println(neworder)

}
