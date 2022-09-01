package main

import (
	"fmt"
	"testing"
)

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func AssertTrue(t *testing.T, got bool, st string) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
	fmt.Println(st)
}

func AssertFalse(t *testing.T, got bool, st string) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
	fmt.Println(st)
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(StackOfInts)

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty(), "Integer1")

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty(), "Integer2")

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		myStackOfInts.Pop()
		myStackOfInts.Pop()
		AssertTrue(t, myStackOfInts.IsEmpty(), "Integer3")
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(StackOfStrings)

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty(), "String1")

		// add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty(), "String2")

		// add another thing, pop it back again
		myStackOfStrings.Push("456")
		myStackOfStrings.Pop()
		myStackOfStrings.Pop()
		AssertTrue(t, myStackOfStrings.IsEmpty(), "String3")
	})
}
