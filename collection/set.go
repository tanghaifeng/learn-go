package collection

import "fmt"

type Set struct {
	m map[interface{}]bool
}

func (s *Set) NewSet() {
	s.m = make(map[interface{}]bool)
}

func (s *Set) Add(e interface{}) {
	s.m[e] = true
}

func (s *Set) Del(e interface{}) {
	if s.m[e] {
		delete(s.m, e)
		fmt.Printf("删除成功")
	}
}

func (s *Set) GetAll() map[interface{}]bool {
	return s.m
}

func (s *Set) ForElement() {
	for k, _ := range s.m {
		fmt.Println(k)
	}
}
