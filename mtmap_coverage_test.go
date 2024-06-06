package mtmap

import (
	"testing"
)

type MyKey[T any] struct{ Key[T] }

func TestCoverage(t *testing.T) {
	k := MyKey[string]{"hello"}
	if k.String() != "hello" {
		t.Fatal("incorrect stringification")
	}
	k.getValue()

	//m := &Map{}
	/*
		if Get(m, MyKey[string]{"hello"}) != "" {
			t.Fatal("incorrect zero value")
		}
	*/
}
