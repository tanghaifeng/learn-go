package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func Concurrency() {

	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup // 跟Java的 CountDownLatch 类似
	wg.Add(2)

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
