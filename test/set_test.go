package test

import (
	"learn-go/collection"
	"testing"
)

func TestSet(t *testing.T) {

	//collection.Map()
	s := collection.Set{}
	s.NewSet()
	s.Add(1)
	s.Add(1)
	s.Add("ssssss")
	s.Add(2.22)
	s.ForElement()
	s.Del(1)
	s.ForElement()
}
