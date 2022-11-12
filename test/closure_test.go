package test

import (
	"log"
	"testing"
)

func TestClosure(t *testing.T) {

	f := closures(1)
	i := f()
	log.Print(i)

}

func closures(i int) func() int {
	return func() int {
		return i
	}
}
