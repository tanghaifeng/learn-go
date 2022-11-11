package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConst(t *testing.T) {
	const age = 18

	fmt.Println(age)

	fmt.Println(strconv.Itoa(age))
}
