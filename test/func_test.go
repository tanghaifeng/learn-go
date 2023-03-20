package test

import (
	"fmt"
	"learn-go/function"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	//s := function.TimeSpent(sleep)
	//s(1)
	//
	//function.VarParams(1, 2, 3, 5)

	function.DeferFunc(func() {
		fmt.Printf("defer func ")
	})
}

func sleep(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
