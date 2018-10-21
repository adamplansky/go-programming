package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var cnt int64

	var wg sync.WaitGroup
	const gs int = 100

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&cnt, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&cnt))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end", cnt)
}
