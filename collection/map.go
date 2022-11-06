package collection

import (
	"fmt"
)

func Map() {
	// literal
	m := map[string]string{
		"name": "Tim",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// make map
	m1 := make(map[string]string)
	m1["addr"] = "shanghai"

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

// 工厂方法
func MapWithFuncValue() {
	m := make(map[int]func(op int) int)
	m[1] = func(op int) int {
		return op
	}

	m[2] = func(op int) int {
		return op * op
	}

	m[3] = func(op int) int {
		return op * op * op
	}
	fmt.Print(m[1](1), m[2](2), m[3](3))
}
