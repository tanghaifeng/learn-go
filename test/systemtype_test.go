package test

import (
	"fmt"
	"learn-go/systemtype"
	"testing"
)

func TestSystemType(t *testing.T) {

	//u := &systemtype.User{Name: "Tim", Age: 30}

	//systemtype.GetUser(u)
	//
	//u.SetInfo(systemtype.User{
	//	Age: 22,
	//})
	//
	//systemtype.GetUser(u)
	//
	//u1 := &systemtype.User{Name: "Joe", Age: 28}
	//systemtype.GetUser(u1)

	//u.Write([]byte("Hello"))
	//fmt.Fprintf(u, "world")
	//
	//var b bytes.Buffer
	//b.Write([]byte("hello"))
	//fmt.Println(&b, "world")

	counter := systemtype.New(1)
	fmt.Printf("%v", counter)
}
