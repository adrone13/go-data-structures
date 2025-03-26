package stack

import (
	"slices"
	"testing"
)

type person struct {
	name string
}

func TestStackStruct(t *testing.T) {
	stack := New[person]()

	people := []person{
		{name: "alex"},
		{name: "ann"},
		{name: "john"},
		{name: "roman"},
		{name: "denis"},
	}
	for _, p := range people {
		stack.Push(p)
	}
	slices.Reverse(people)
	for _, expected := range people {
		received := stack.Pop()
		if expected != received {
			t.Errorf("Expected popped item to be %+v but got %+v", expected, received)
		}
	}
}

func TestStackInt(t *testing.T) {
	stack := New[int]()

	nums := []int{1, 2, 3, 4, 5}
	for _, p := range nums {
		stack.Push(p)
	}
	slices.Reverse(nums)
	for _, expected := range nums {
		received := stack.Pop()
		if expected != received {
			t.Errorf("Expected popped item to be %d but got %d", expected, received)
		}
	}
}
