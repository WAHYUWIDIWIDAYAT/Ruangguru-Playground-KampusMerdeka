package perusahaan

import "fmt"

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	var sum int
	for _, v := range c.Subordinate {
		sum = sum + v.TotalDivisonSalary()
	}
	fmt.Println(c.GetSalary())
	return c.GetSalary() + sum
}