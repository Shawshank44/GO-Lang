package main

import "fmt"

func main() {
	// Declaration and syntax:
	// var mapVariable map[keytype] valuetype
	// mapVariable = make(map[keytype]valuetype)

	// using a Map literal :
	// mapVariable = map[keytype]valuetype{
	// 	key1 : value1,
	// 	key2 : value2,
	// }

	// ages := make(map[string]int)

	// ages["Jack"] = 26
	// ages["John"] = 35
	// ages["Bob"] = 40
	// ages["Cindy"] = 56

	// fmt.Println(ages)

	// Accessing, deleting and modifying :
	// ages["Bob"] = 39          // modifying
	// fmt.Println(ages["John"]) // accessing
	// delete(ages, "Bob") // deleting the key
	// fmt.Println(ages)

	//empty a map
	// clear(ages)
	// fmt.Println(ages)

	//Bool value using maps
	// value, booleans := ages["Bob"]
	// fmt.Println(value, booleans)

	// Iteration
	// mp := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// for k, v := range mp {
	// 	fmt.Printf("Key : %s ; Value : %d \n", k, v)
	// }

	// fmt.Println(len(mp)) // returns length of the map

	// Nested maps :
	electronics := map[string]string{
		"Laptop":          "$500",
		"Washing machine": "$800",
		"Mobile":          "$200",
	}

	Groceries := map[string]string{
		"Apple":    "$3",
		"Greens":   "$10",
		"Biscuits": "$0.10",
	}

	Medicines := map[string]string{
		"Paracetmol": "$1.5",
		"Eldopher":   "$0.12",
		"Horse-P":    "$10",
	}

	shop := make(map[string]map[string]string)
	shop["Electronics"] = electronics
	shop["Groceries"] = Groceries
	shop["Medicines"] = Medicines

	for category, items := range shop {
		fmt.Println("Category:", category)
		for item, price := range items {
			fmt.Printf("  %s: %s\n", item, price)
		}
	}

}
