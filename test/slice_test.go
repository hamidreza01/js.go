package test

import (
	"testing"

	"github.com/hamidreza01/js.go/slice"
)

func TestPush(t *testing.T) {
	var s slice.Slice[int]
	s.Push(1, 2, 3)
	s.Push(4, 5)
	if len(s) < 5 {
		t.Errorf("Push method is not working properly, out : %#v", s)
	}
}

func TestMap(t *testing.T) {
	s := slice.Slice[int]{
		5, 10,
	}
	s.Map(func(index int, value int) int {
		return value * 2
	})
	if s[0] != 10 && s[1] != 20 {
		t.Errorf("Map method is not working properly, out : %#v", s)
	}
}

func TestFilter(t *testing.T) {
	s := slice.Slice[int]{
		1,
		2,
		10,
		30,
		40,
	}
	s.Filter(func(index int, value int) bool {
		return value >= 10
	})
	if len(s) > 3 {
		t.Errorf("Filter method is not working properly, out : %#v", s)
	}
}

func TestForEach(t *testing.T) {
	var i int8
	s := slice.Slice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s.ForEach(func(index int, value int) {
		i++
	})
	if i != 9 {
		t.Errorf("ForEach method is not working properly")
	}
}

func TestConcat(t *testing.T) {
	s := slice.Slice[int]{1, 2, 3, 4, 5}
	s2 := slice.Slice[int]{5, 7, 8, 9}
	s.Concat(s2)
	if len(s) < 9 {
		t.Errorf("Filter method is not working properly, out : %#v", s)
	}
}

func TestEvery(t *testing.T) {
	s := slice.Slice[int]{1, 2, 3, 4, 6, 7, 8, 9, 10}
	b := s.Every(func(index int, value int) bool {
		return value > 10
	})
	if b {
		t.Errorf("Every method is not working properly, out : %t", b)
	}
}

func TestRemove(t *testing.T) {
	s := slice.Slice[int]{1, 2, 3}
	s.Remove(2)
	if len(s) > 2 {
		t.Errorf("Remove method is not working properly, out : %#v", s)
	}
}

func TestFind(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	s := slice.Slice[person]{
		person{
			name: "hamid",
			age:  21,
		},
		person{
			name: "mahdi",
			age:  16,
		},
	}
	v, err := s.Find(func(i int, v person) bool {
		return v.name == "mahdi"
	})
	if err != nil || v.name != "mahdi" {
		t.Errorf("Find method is not working properly, out : %#v", v)
	}
}
func TestFindIndex(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	s := slice.Slice[person]{
		person{
			name: "mahdi",
			age:  16,
		},
		person{
			name: "hamid",
			age:  21,
		},
	}
	index := s.FindIndex(func(i int, v person) bool {
		return v.name == "hamid"
	})
	if s[index].name != s[1].name {
		t.Errorf("FindIndex method is not working properly, out : %#v", s)
	}
}

func TestSome(t *testing.T) {
	s := slice.Slice[int]{1, 2, 3, 10, 6, 7, 8, 9}
	if !s.Some(func(index int, value int) bool {
		return value == 10
	}) {
		t.Errorf("Some method is not working properly")
	}
	if s.Some(func(index int, value int) bool {
		return value > 100
	}) {
		t.Errorf("Some method is not working properly")
	}
}

func TestIncludes(t *testing.T) {
	s := slice.Slice[int]{1, 2, 3, 10, 6, 7, 8, 9}
	if !slice.Includes(s, 10) {
		t.Errorf("Includes method is not working properly")
	}
	type person struct {
		name string
		age  int
	}
	s2 := slice.Slice[person]{
		person{
			name: "mohammad",
			age:  37,
		},
		person{
			name: "navid",
			age:  18,
		},
	}
	if !slice.Includes(s2, person{name: "mohammad", age: 37}) {
		t.Errorf("Includes method is not working properly")
	}
}

func TestJoin(t *testing.T) {
	s := slice.Slice[string]{
		"Hello",
		"World",
	}
	s2 := slice.Join(s, " ")
	if s2 != "Hello World" {
		t.Errorf("`slice.Slice.Join` method is not working properly, out : %v", s2)
	}
}

func TestReduce(t *testing.T) {
	s := slice.Slice[int]{
		5, 5, 5, 5,
	}
	i := s.Reduce(func(total, current int) int {
		return total + current
	})
	if i != 20 {
		t.Errorf("`slice.Slice.Reduce` method is not working properly, out : %v", i)
	}
}

func TestIndexOf(t *testing.T) {
	s := slice.Slice[string]{
		"hamid", "mahdi",
	}
	if slice.IndexOf(s, "mahdi") != 1 && slice.IndexOf(s, "hossein") != -1 {
		t.Errorf("`slice.Slice.IndexOf` method is not working properly")
	}
}
