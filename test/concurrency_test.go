package test

import (
	"fmt"
	"learn/concurrency"
	"testing"
)

func TestConcurrency(t *testing.T) {
	//concurrency.Concurrency()
	//
	//concurrency.Race()

	//concurrency.Atomic()
	//concurrency.Channel()
	//r := concurrency.New(time.Second * 10)
	//r.Add(a)
	//r.Add(b)
	//r.Add(c)
	//r.Start()
	//concurrency.ProducerAndReceiver()
	s1 := concurrency.GetInstance()
	s2 := concurrency.GetInstance()
	if s1 == s2 {
		fmt.Println("===")
	}
}

func a(c int) {

	fmt.Println("funca")
}

func b(c int) {
	select {}
	fmt.Println("funcb")
}

func c(c int) {
	fmt.Println("funcb")
}
