package test

import "testing"

func TestValue(t *testing.T) {

	t.Logf("%d", 1+2)

	t.Logf("%s", "1"+"2")

	t.Logf("%f", 7.0/3.0)

	t.Log(true || false)
}
