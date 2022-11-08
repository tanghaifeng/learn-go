package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {

	// 连接字符串
	var builder strings.Builder
	for i := 0; i < 100; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	fmt.Println(builder.String())

	u := strings.ToUpper("a")
	fmt.Println(u)

	joins := []string{"Tim", "Age"}

	fmt.Println(strings.Join(joins, "-"))

}
