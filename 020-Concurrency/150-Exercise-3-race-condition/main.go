package main

import (
	"fmt"
	"sync"
)

func main() {

	cnt := 0

	var wg sync.WaitGroup
	const gs int = 100
	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := cnt
			v++
			// runtime.Gosched()
			cnt = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end", cnt)
}
