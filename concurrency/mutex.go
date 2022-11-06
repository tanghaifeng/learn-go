package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	mutex sync.Mutex
)

func Sync() {
	wg.Add(2)

	go intCounter1()

	go intCounter1()

	wg.Wait()
	fmt.Println(counter)

}

func intCounter1() {

	defer wg.Done()

	for count := 0; count < 2; count++ {
		defer func() {
			mutex.Unlock()
		}()
		mutex.Lock()
		{
			value := counter

			runtime.Gosched()

			value++
			counter = value
		}
	}

}
