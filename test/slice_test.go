package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	//collection.Slice()

	book := make([]string, 1)
	book = append(book, "童话故事", "老人与海")
	test(book)
	for _, v := range book {
		fmt.Println(v)
	}

}

func test(b []string) {

	for _, v := range b {
		fmt.Println(v)
	}
	b[1] = "老子"
}
