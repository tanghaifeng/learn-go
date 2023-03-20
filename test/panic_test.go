package test

import (
	"fmt"
	panic_demo "learn-go/panic"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
		}
	}()
	panic_demo.Panic()
}
