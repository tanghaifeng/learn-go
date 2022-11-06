package panic_demo

import (
	"errors"
	"fmt"
)

func Panic() {
	fmt.Println("运行正常")

	panic(errors.New("panic"))
}
