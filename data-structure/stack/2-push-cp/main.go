package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

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
