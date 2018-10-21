package main

import (
	"fmt"
)

type S struct {
	i int
}

func (p S) Get() int {
	fmt.Printf("Get()\t %T %v\n", p, p)
	return p.i
}
func Get() int {
	fmt.Println("GET() without references")
	return 1
}

type I interface {
	Get() int
}

func f(p I) {
	fmt.Printf("INTERFACE f(p I) \t%T %v\n", p, p)
	// fmt.Println(p.Get())
	// p.Put(1)
}

func main() {
	s := S{5}
	fmt.Println(s)
	// f(&s)
	// f(s)
	a := (&s).Get()
	fmt.Println(a)
}

// package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// type circle struct {
// 	radius float64
// }
//
// type shape interface {
// 	area() float64
// 	area1() float64
// }
//
// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
//
// func (c circle) area1() float64 {
// 	return math.Pi * c.radius * c.radius
// }
//
// func info(s shape) {
// 	fmt.Println("area", s.area())
// }
// func info1(s shape) {
// 	fmt.Println("area", s.area1())
// }
//
// func main() {
// 	c := circle{5}
// 	// why is it not working??
// 	// The method set of a type determines the INTERFACES that the type implements.....
// 	info(&c)
// 	info1(&c)
// 	// fmt.Println(c.area())
// }
