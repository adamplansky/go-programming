// package main
//
// import "fmt"
//
// func sum(s []int, c chan int) {
// 	fmt.Println(s)
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }
//
// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}
//
// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c
//
// 	fmt.Println(x, y, x+y)
// }

package main

import "fmt"

func fact(s []int, c chan int) {
	f := 1
	fmt.Println(s)
	for _, v := range s {
		f *= v
	}
	c <- f
}
func factorial(n int) {
	c := make(chan int)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	go fact(s[:len(s)/2], c)
	go fact(s[len(s)/2:], c)
	fmt.Println(c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x*y)
}
func main() {
	// ch := make(chan int)
	// go func() {
	// 	ch <- doSomething(5)
	// }()
	// fmt.Println(<-ch)
	factorial(10)
	//fmt.Println(x)
}
