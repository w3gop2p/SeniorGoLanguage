package main

import "fmt"

// Generic types, in this case structs
type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

// Generic functions
func foo[T any, B any](val T, x B) {
	fmt.Println(val)
	fmt.Println(x)
}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("bar", 2)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.99)
	m2.Insert(2, 100.333)

	// generic functions
	foo[string, int]("agg", 12)
}
