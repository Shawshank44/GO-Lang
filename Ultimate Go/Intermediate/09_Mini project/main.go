package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	Name        string
	Age         int
	PhoneNumber string
}

type Products struct {
	Product []string
}

type Order struct {
	Id           string
	ProductName  map[string]int
	Quantity     int
	CustomerName string
	Subtotal     float32
}

type Delivery struct {
	Id           string
	CustomerName string
	Status       string
}

// Seed random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random order ID
func generateOrderID() string {
	return fmt.Sprintf("ORDER-%d", rand.Intn(1000000))
}

func (c *Customer) Createcustomer(Name string, Age int, PhoneNumber string) {
	c.Name = Name
	c.Age = Age
	c.PhoneNumber = PhoneNumber
}

func (p *Products) CreateProduct(products ...string) {
	p.Product = products
}

// Check if all requested products exist in Products list
func (p *Products) AreProductsAvailable(requested map[string]int) bool {
	for productName := range requested {
		found := false
		for _, existingProduct := range p.Product {
			if existingProduct == productName {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func NewOrder(o *Order, c *Customer, requestedProducts map[string]int, pricePerProduct map[string]float32, p *Products) bool {
	if !p.AreProductsAvailable(requestedProducts) {
		fmt.Println("❌ Error: One or more requested products are not available.")
		return false
	}

	o.Id = generateOrderID()
	o.ProductName = requestedProducts
	o.CustomerName = c.Name

	var subtotal float32
	var totalQty int

	for product, qty := range requestedProducts {
		price := pricePerProduct[product]
		subtotal += float32(qty) * price
		totalQty += qty
	}

	o.Subtotal = subtotal
	o.Quantity = totalQty

	return true

}

func main() {
	var customer Customer
	customer.Createcustomer("Alice", 28, "555-1234")

	var products Products
	products.CreateProduct("Laptop", "Phone", "Headphones")

	pricePerProduct := map[string]float32{
		"Laptop":     1000.0,
		"Phone":      700.0,
		"Headphones": 150.0,
	}

	// Example of a valid product request
	productRequest := map[string]int{
		"Phone":      2,
		"Headphones": 1,
	}

	var order Order
	success := NewOrder(&order, &customer, productRequest, pricePerProduct, &products)

	if success {
		fmt.Println("✅ Order created successfully:")
		fmt.Println("Order ID:", order.Id)
		fmt.Println("Customer:", order.CustomerName)
		fmt.Println("Subtotal:", order.Subtotal)

		delivery := Delivery{
			Id:           order.Id,
			CustomerName: customer.Name,
			Status:       "Processing",
		}

		fmt.Println("Delivery Info:", delivery)
	}

}
