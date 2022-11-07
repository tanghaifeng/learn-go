package test

import (
	"fmt"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		printXX()
	}
	b.StopTimer()
}

func printXX() {
	fmt.Println("xx")
}
