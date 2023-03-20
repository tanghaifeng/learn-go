package test

import (
	"fmt"
	error_type "learn-go/error"
	"testing"
)

func TestError(t *testing.T) {
	var (
		i   int
		err error
	)
	if i, err = error_type.ErrorFun(10); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)

}
