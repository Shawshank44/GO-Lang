package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func age(name string, age int) (string, int) {
	return name, age
}

func getprogramlanguages() (string, string, string) {
	return "Golang", "Javascript", "C/C++"
}

func processData(data int, callback func(int)) {
	fmt.Println("Processing data : ", data)
	callback(data)
}

func printResult(result int) {
	fmt.Println("Processed result:", result*2)
}

func main() {

	fmt.Println(add(4, 5))

	fmt.Println(age("John", 33))

	lang1, lang2, lang3 := getprogramlanguages() // note if dont need to use any return type mark _

	fmt.Println(lang1, lang2, lang3)

	processData(10, printResult)

}
