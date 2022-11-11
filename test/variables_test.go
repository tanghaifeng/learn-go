package test

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {
	var a = "a"
	fmt.Println(a)

	b, c := 1, 2
	fmt.Println(b, c)
}
