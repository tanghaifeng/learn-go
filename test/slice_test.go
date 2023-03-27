package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	//collection.Slice()

	//book := make([]string, 1)
	//book = append(book, "童话故事", "老人与海")
	//test(book)
	//for _, v := range book {
	//	fmt.Println(v)
	//}
	testNilSlice()
}

func test(b []string) {

	for _, v := range b {
		fmt.Println(v)
	}
	b[1] = "老子"
}

func testNilSlice() {
	// 空的切片
	//s := make([]string, 0)
	// nil 切片
	var s []string
	s = append(s, "11", "2222")
	fmt.Println(strings.Join(s, ","))

}
