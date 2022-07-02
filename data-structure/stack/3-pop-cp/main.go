package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")
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

func (s *Stack) Pop() (int, error) {
	if s.Top == -1 {
		return 0, ErrStackUnderflow
	}
	Elemen := s.Data[s.Top]
	s.Top--
	s.Data = s.Data[:s.Top+1]
	return Elemen, nil
}
