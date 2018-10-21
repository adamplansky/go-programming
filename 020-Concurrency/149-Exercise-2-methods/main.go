package main

import "fmt"

type person struct {
	Name string
}

func (p *person) speak() string {
	return fmt.Sprint("I am speaking ", p.Name)
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p := person{"Adam"}
	saySomething(&p)

	// cannot use
	//saySomething( p)

	fmt.Println(p.speak())
	fmt.Println((&p).speak())
}
