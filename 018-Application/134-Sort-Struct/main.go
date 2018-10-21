package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("<%v %d>", p.Name, p.Age)
}

//https://godoc.org/sort#Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//https://godoc.org/sort#Interface
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	persons := []Person{
		{"Adam", 26},
		{"Jitka", 45},
		{"Alzbeta", 17},
	}
	sort.Sort(ByName(persons))
	fmt.Println(persons)
	sort.Sort(ByAge(persons))
	fmt.Println(persons)
}
