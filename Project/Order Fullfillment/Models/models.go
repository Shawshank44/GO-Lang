package models

import (
	"fmt"
	"time"
)

type Order struct {
	ID       int
	Customer string
	Item     string
	Quantity int
}

func (o *Order) Fulfill() {
	fmt.Printf("Fulfilling Order #%d for %s : %dx %s \n", o.ID, o.Customer, o.Quantity, o.Item)
	time.Sleep(2 * time.Second)
	fmt.Printf("Order #%d fulfilled \n", o.ID)
}
