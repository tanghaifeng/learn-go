package test

import (
	"fmt"
	"testing"
)

/*
*
递归测试
*/
func TestRecursion(t *testing.T) {
	recursion(1)
}

func recursion(i int) int {
	if i > 100 {
		fmt.Println("break recursion")
		return i
	}
	i++
	fmt.Println("go recursion")
	return recursion(i)
}
