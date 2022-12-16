package network

import (
	"fmt"
	"net/http"
)

func Http() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := &http.Cookie{}
		c.Name = "tim"
		c.Value = "TestCookie"
		http.SetCookie(writer, c)
		fmt.Fprintln(writer, "Hello Http")
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello")
	})

	http.ListenAndServe(":80", nil)
}
