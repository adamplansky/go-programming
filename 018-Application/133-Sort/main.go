package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{5, 1, 2, 100, 23, 55, 19, 28, 3}
	fmt.Println("not sorted", n)
	sort.Ints(n)
	fmt.Println("sorted", n)

	a := []int{5, 1, 2, 100, 23, 55, 19, 28, 3}
	a.sort()
	fmt.Println("sorted", a)
}
