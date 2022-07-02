package laptop

type Sleeping struct {
	Laptop *Laptop
}

func (s Sleeping) Press() {
	s.Laptop.CurrentState = "On"
	s.Laptop.ChangeState(On{s.Laptop})
}

func (s Sleeping) CanTurnOnLaptop() bool {
	return true
}

func (s Sleeping) Sleep() {
	s.Laptop.CurrentState = "Sleeping"
	s.Laptop.ChangeState(Sleeping{s.Laptop})
}