package concurrency

import (
	"fmt"
	"math/rand"
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
