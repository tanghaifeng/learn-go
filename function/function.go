package function

import (
	"fmt"
	"time"
)

// 函数是一等公民
func TimeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Printf("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func VarParams(op ...int) {
	for v := range op {
		fmt.Println(v)
	}
}

func DeferFunc(d func()) {
	defer d()
	fmt.Println("先执行")
}
