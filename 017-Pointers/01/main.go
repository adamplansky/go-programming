package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	(*p).name = "changed"
}

//not thread safe but better for single thread
func (p *person) changeMePointer() {
	(*p).name = "changed pointer"
}

func main() {
	p := person{
		name: "adam",
	}
	fmt.Printf("%T \n", p)
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
	p.changeMePointer()
	fmt.Println(p)
}
