package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func Race() {
	wg.Add(2)
	go intCounter()
	go intCounter()

	wg.Wait()
	fmt.Println(counter)

}

func intCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {

		value := counter

		runtime.Gosched()

		value++
		counter = value
	}
}
