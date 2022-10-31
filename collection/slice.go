package collection

import "fmt"

func Slice() {
	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(s[1])

	s1 := make([]int, 10)
	for _, v := range s1 {
		fmt.Println(v)
	}

	fmt.Println(s[:])
	fmt.Println(s[0:1])

	newSlice := s[0:1]
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

}
