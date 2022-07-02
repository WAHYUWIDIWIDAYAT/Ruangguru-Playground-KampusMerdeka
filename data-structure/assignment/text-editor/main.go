package main

import (

	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
		
	}
}

func (te *TextEditor) Write(ch rune) {
	te.RedoStack.SetToEmpty()
	te.UndoStack.Push(ch)
}

func (te *TextEditor) Read() []rune {
	// TODO: answer 
	var result []rune
	for te.UndoStack.Top > -1 {
		char, _ := te.UndoStack.Pop()
		result = append(result, char)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	
	}
	return result
}

func (te *TextEditor) Undo() {
	// TODO: answer here
	char, err := te.UndoStack.Peek()
	if err != nil {
		return
	}
	te.UndoStack.Pop()
	te.RedoStack.Push(char)
}

func (te *TextEditor) Redo() {
	char, err := te.RedoStack.Peek()
	if err != nil {
		return
	}
	te.RedoStack.Pop()
	te.UndoStack.Push(char)
}
