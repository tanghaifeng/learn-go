package concurrency

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var (
	counter1 int64
)

func Atomic() {
	wg.Add(2)

	go intCounter2(1)

	go intCounter2(2)

	wg.Wait()
	fmt.Println(counter1)

}

func intCounter2(id int) {

	defer wg.Done()

	for count := 0; count < 2; count++ {

		atomic.AddInt64(&counter1, 1)

		runtime.Gosched()

	}

}
