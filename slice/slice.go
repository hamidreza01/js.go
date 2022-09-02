package slice

import (
	"errors"
	"strings"
)

type Slice[T any] []T

func (s *Slice[T]) Push(v ...T) *Slice[T] {
	*s = append(*s, v...)
	return s
}

func (s *Slice[T]) Map(f func(index int, value T) T) *Slice[T] {
	l := len(*s)
	for i := 0; i < l; i++ {
		(*s)[i] = f(i, (*s)[i])
	}
	return s
}

func (s *Slice[T]) Filter(f func(index int, value T) bool) *Slice[T] {
	for i, v := range *s {
		if f(i, v) {
			*s = append((*s)[:len(*s)-1], (*s)[len(*s):]...)
		}
	}
	return s
}

func (s *Slice[T]) ForEach(f func(index int, value T)) *Slice[T] {
	for i, v := range *s {
		f(i, v)
	}
	return s
}

func (s *Slice[T]) Concat(v []T) *Slice[T] {
	*s = append(*s, v...)
	return s
}

func (s *Slice[T]) Every(f func(index int, value T) bool) bool {
	for i, v := range *s {
		if !f(i, v) {
			return false
		}
	}
	return true
}

func (s *Slice[T]) Some(f func(index int, value T) bool) bool {
	for i, v := range *s {
		if f(i, v) {
			return true
		}
	}
	return false
}

func (s *Slice[T]) Remove(index int) *Slice[T] {
	*s = append((*s)[:index], (*s)[index+1:]...)
	return s
}

func (s *Slice[T]) Reverse() *Slice[T] {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return s
}

func (s *Slice[T]) Slice(i1 int, i2 int) Slice[T] {
	return (*s)[i1:i2]
}

func (s *Slice[T]) Fill(value T) *Slice[T] {
	for i := range *s {
		(*s)[i] = value
	}
	return s
}

func (s *Slice[T]) Find(f func(index int, value T) bool) (value T, err error) {
	for i, v := range *s {
		if f(i, v) {
			return v, nil
		}
	}
	return *new(T), errors.New("the desired value could not be found")
}

func (s *Slice[T]) FindIndex(f func(index int, value T) bool) int {
	for i, v := range *s {
		if f(i, v) {
			return i
		}
	}
	return -1
}

func (s *Slice[T]) Reduce(f func(total T, current T) T) T {
	var total T
	for _, v := range *s {
		total = f(total, v)
	}
	return total
}

func Join(s Slice[string], sep string) string {
	return strings.Join(s, sep)
}

func Includes[T comparable](s Slice[T], value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](s Slice[T], value T) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}
	return -1
}
