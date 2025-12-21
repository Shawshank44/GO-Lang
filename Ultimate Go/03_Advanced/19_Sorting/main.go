package main

import (
	"fmt"
	"sort"
)

// sorting by Function
type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type PersonSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *PersonSorter) Len() int {
	return len(s.people)
}

func (s *PersonSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *PersonSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {
	// numbers := []int{5, 10, 6, 2, 1, 3, 4, 7, 8, 9}
	// sort.Ints(numbers) // sort numbers
	// fmt.Println(numbers)

	// stringSLC := []string{"John", "Jack", "Simon", "Steve", "victor", "Walter"}
	// sort.Strings(stringSLC) // sort strings
	// fmt.Println(stringSLC)

	// Sorting by functions :
	People := []Person{
		{"Alice", 30},
		{"Bobby", 35},
		{"Anna", 25},
	}

	// sort.Sort(ByAge(People))
	// sort.Sort(ByName(People))
	AsenAge := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	Name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	DesenAge := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	LenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	By(AsenAge).Sort(People)
	fmt.Println("Sorted by age (asending) : ", People) // Small to Big
	By(DesenAge).Sort(People)
	fmt.Println("Sorted by age (Desending) : ", People) // Big to Small
	By(Name).Sort(People)
	fmt.Println("Sorted by Name : ", People)
	By(LenName).Sort(People)
	fmt.Println("Sorted by LengthofName : ", People)

	// Sort.Slice Sorting by function
	stringslice := []string{"Banana", "apple", "cherry", "grapes", "guava", "rippler"}
	sort.Slice(stringslice, func(i, j int) bool {
		return stringslice[i][len(stringslice[i])-1] < stringslice[j][len(stringslice[j])-1]
	})
	fmt.Println("Sorted by last character : ", stringslice)
}
