package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	value := s[i].Correct*4 - s[i].Wrong
	value2 := s[j].Correct*4 - s[j].Wrong

	if value > value2 {
		return true
	} else if value == value2 {
		if s[i].Correct > s[j].Correct {
			return true
		} else if s[i].Correct == s[j].Correct {
			sort.Strings([]string{s[i].Name, s[j].Name})
			return s[i].Name < s[j].Name
		}
	}
	return false // TODO: replace this
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})
	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}