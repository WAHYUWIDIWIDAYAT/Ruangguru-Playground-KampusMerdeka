package coffe

type Mocha struct {
	//beginanswer
	Coffe Coffe
	//endanswer
}

func (m Mocha) GetCost() float64 {
	//beginanswer
	return m.Coffe.GetCost() + 1.00
	//endanswer return 0
}

func (m Mocha) GetDescription() string {
	//beginanswer
	return m.Coffe.GetDescription() + ", Mocha"
	//endanswer return ""
}

type Whipcream struct {
	//beginanswer
	Coffe Coffe
	//endanswer
}

func (w Whipcream) GetCost() float64 {
	//beginanswer
	return w.Coffe.GetCost() + 0.10
	//endanswer return 0
}

func (w Whipcream) GetDescription() string {
	//beginanswer
	return w.Coffe.GetDescription() + ", Whipcream"
	//endanswer return ""
}