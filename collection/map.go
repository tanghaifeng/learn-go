package collection

import "fmt"

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
