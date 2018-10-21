package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func printSmth(text string) {
	fmt.Print(text)
	wg.Done()
}

func main() {
	const gs int = 10000
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		s := strconv.Itoa(i)

		go printSmth("t " + s)
	}

	wg.Wait()
	fmt.Println("end")
}
