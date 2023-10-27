package objectpool

import (
	"slices"
)

// ObjectPool stores a collection of object pointers with utility functions to manipulate the pool.
//
// Important note: elements in an ObjectPool will NOT stay in the same order.
type ObjectPool[T comparable] struct {
	Objects []T
}

type ObjectNotFoundError struct{}

func NewObjectPool[T comparable]() *ObjectPool[T] {
	return &ObjectPool[T]{
		Objects: make([]T, 0),
	}
}

func (pool *ObjectPool[T]) Add(obj T) {
	pool.Objects = append(pool.Objects, obj)
}

func (pool *ObjectPool[T]) Remove(index int) {
	sliceLen := len(pool.Objects)
	pool.Objects[index] = pool.Objects[sliceLen-1]
	//pool.Objects[sliceLen-1] = nil
	pool.Objects = pool.Objects[:sliceLen-1]
}

func (pool *ObjectPool[T]) FindAndRemove(obj T) error {
	index, err := pool.Find(obj)
	if err != nil {
		return err
	}
	pool.Remove(index)
	return nil
}

func (pool *ObjectPool[T]) Find(obj T) (int, error) {
	for i, object := range pool.Objects {
		if object == obj {
			return i, nil
		}
	}
	return 0, &ObjectNotFoundError{}
}

func (pool *ObjectPool[T]) Contains(obj T) bool {
	return slices.Contains(pool.Objects, obj)
}

func (e *ObjectNotFoundError) Error() string {
	return "Object Not Found"
}
