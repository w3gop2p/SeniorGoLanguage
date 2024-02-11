package main

import "fmt"

type Putter interface {
	Put(id int, val any) error
	fmt.Stringer
}

type Storage interface {
	Putter
	Get(id int) (int, error)
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (int, error) {
	fmt.Println("Get implementation of the FooStorage type")
	return id, nil
}

func (s *FooStorage) Put(id int, val any) error {
	fmt.Println("Put implementation of the FooStorage type")
	return nil
}

func (s *FooStorage) String() string { return "" }

type simplePutter struct{}

func (s *simplePutter) Get(id int) (int, error) {
	fmt.Println("Get implementation of the SimplePutter type")
	return id, nil
}

func (s *simplePutter) Put(id int, val any) error {
	fmt.Println("Put implementation of the SimplePutter type")
	return nil
}

func (s *simplePutter) String() string { return "" }

type Server struct {
	store Storage
}

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	s.store.Get(1)
	s.store.Put(1, "foo")
	updateValue(3, "foo", s.store)

	// simplePutter
	sputter := &simplePutter{}
	updateValue(1, "foo", sputter)
}
