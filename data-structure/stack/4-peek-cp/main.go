package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")
var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	Top int
	Data []int
	Size int

}

func NewStack(size int) Stack {
	return Stack{
		Top: -1,
		Data: []int{},
		Size: size,
	}
}

func (s *Stack) Push(Elemen int) error {
	if s.Top == s.Size-1 {
		return ErrStackOverflow
	}
	s.Top++
	s.Data = append(s.Data, Elemen)
	return nil
}

func (s *Stack) Peek() (int, error) {
	if s.Top == -1 {
		return 0, ErrEmptyStack
	}
	Elemen := s.Data[s.Top]
	return Elemen, nil
}
