package mtmap_test

import (
	"fmt"

	"github.com/thejerf/mtmap"
)

type MyTypedKey[T any] struct{ mtmap.Key[T] }

type MyBoolKey struct{ mtmap.Key[bool] }

func Example_mtMap() {
	m := &mtmap.Map{}

	// Set some values
	mtmap.Set(m, MyTypedKey[string]{"hello"}, "world")
	// Note keys namespaced by type
	mtmap.Set(m, MyTypedKey[int]{"hello"}, 25)
	// key with hardcoded value easier to use
	mtmap.Set(m, MyBoolKey{"hello"}, true)

	// You can now retrieve values in a type-safe manner:
	world := mtmap.Get(m, MyTypedKey[string]{"hello"})
	fmt.Println("MyTypedKey[string]{\"hello\"}:", world)

	isHello := mtmap.Get(m, MyBoolKey{"hello"})
	fmt.Println("bool hello:", isHello)

	_, haveKey := mtmap.Get2(m, MyTypedKey[string]{"no_such"})
	fmt.Println("no_such key:", haveKey)

	// note this is a compiler error; you must use a key type.
	// mtmap.Set(m, "not a key", "hello")

	// Output:
	// MyTypedKey[string]{"hello"}: world
	// bool hello: true
	// no_such key: false
}
