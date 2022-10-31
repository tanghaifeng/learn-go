package collection

import "fmt"

func Array() {

	// array declaration
	// zero value ""  interface{} nil  int 0
	var a [1]string

	// zero value
	fmt.Println(a[0])
}

func ForeachArray() {
	a := [...]int{2, 8, 10}
	for _, v := range a {
		fmt.Println(v)
	}
}
