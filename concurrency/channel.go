package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
)

func Channel() {
	c := make(chan int)
	wg.Add(2)
	go play("Tim", c)
	go play("Ben", c)
	c <- 1
	wg.Wait()
}

func play(name string, count chan int) {
	defer wg.Done()

	for {
		ball, ok := <-count
		if !ok {
			fmt.Printf("Play %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Play %s won\n", name)
			close(count)
			return
		}
		fmt.Printf("Play %s Hit %d \n", name, ball)
		ball++
		count <- ball
	}
}

// 生产者消费者
func ProducerAndReceiver() {
	var wait sync.WaitGroup
	ch := make(chan int)
	wait.Add(1)
	producer(ch, &wait)

	wait.Add(1)
	receiver(ch, &wait)
	wait.Wait()

}

func producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
