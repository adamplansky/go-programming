package main

import (
	"fmt"
	"sort"
)

type Mobile struct {
	Brand string
	Price int
}

// ByPrice implements sort.Interface for []Mobile based on
// the Price field.
type ByPrice []Mobile

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

func main() {
	mobile1 := []Mobile{
		{"Sony", 952},
		{"Nokia", 468},
		{"Apple", 1219},
		{"Samsung", 1045},
	}
	fmt.Println("\nFound mobile1 price is sorted :", sort.IsSorted(ByPrice(mobile1))) // false
	sort.Sort(ByPrice(mobile1))
	fmt.Println(mobile1)
	mobile2 := []Mobile{
		{"Sony", 452},
		{"Nokia", 768},
		{"Apple", 919},
		{"Samsung", 1045},
	}
	fmt.Println("\nFound mobile2 price is sorted :", sort.IsSorted(ByPrice(mobile2))) // true
}
