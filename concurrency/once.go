package concurrency

import (
	"fmt"
	"sync"
)

type Singleton struct {
	m map[string]interface{}
}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("执行一次")
		singleton = &Singleton{}
	})
	return singleton
}

func (s *Singleton) Add(key string, obj interface{}) {
	s.m[key] = obj
}

func (s *Singleton) Del(key string) {
	if s.m[key] != nil {
		delete(s.m, key)
	}
}

func (s *Singleton) Get(key string) interface{} {
	if obj, ok := s.m[key]; ok {
		return obj
	}
	return nil
}
