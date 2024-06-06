package mtmap

import "fmt"

// A Map allows an underlying map[any]any to be accessed in a
// type-safe manner, where the key determines the type of the value
// fully and no type checking is required.
//
// An empty Map can be contructed with &mtmap.Map{}.
//
// Maps have no thread safety implemented in this module.
type Map struct {
	data map[any]any
}

type key[T any] interface {
	getValue() T
	fmt.Stringer
}

// Get2 retrieves the value out of the Map like a two-element access
// of a normal Go map.
func Get2[V any](m *Map, key key[V]) (V, bool) {
	var zero V
	if m.data == nil {
		return zero, false
	}
	val, exists := m.data[key]
	if !exists {
		return zero, false
	}
	return val.(V), true
}

// Get retrieves the value out of a map like a one-element access of a
// normal Go map, returning the zero value of the correct type if the
// key is not present.
func Get[V any](m *Map, key key[V]) V {
	v, _ := Get2(m, key)
	return v
}

// Set will set the value of the given key.
func Set[V any](m *Map, key key[V], val V) {
	if m.data == nil {
		m.data = map[any]any{}
	}

	m.data[key] = val
}

// Key is the key type for setting a value in a TypedMap. The generic
// type indicates not the type of the key, but the type of the VALUE
// stored under that key, e.g., Key[bool] is STORING a bool. The key
// is a string regardless of the type it stores.
//
// You should generally not use this directly, but derive a type from
// it in the package storing a value. This is similar to how values
// stored in a context.Context should have their own type, to
// guarantee no collisions with other possible keys of the same name.
//
// You can either create a generic key in your own package:
//
//	type MyKeyType[T any] struct { mtmap.Key[T] }
//
// which can be used with:
//
//	key := MyKeyType[int]{"the_key"}
//
// to store an int with that key, and then other types with similar
// generic specifications, or you can hard-code the type of the value
// in the key:
//
//	type BoolKey struct { mtmap.Key[bool] }
//
// which simplifies the usage of the key but limits it to one type.
//
// Not exporting a given key can also allow you to require other users
// to access a key's value only through specialized functions or
// methods in a given package, allowing constraints and guarantees to
// be modularly provided on a per-key-type basis.
type Key[Value any] string

// this is not called, but it implements the type restriction
func (k Key[V]) getValue() (val V) { return }

func (k Key[V]) String() string {
	return string(k)
}
