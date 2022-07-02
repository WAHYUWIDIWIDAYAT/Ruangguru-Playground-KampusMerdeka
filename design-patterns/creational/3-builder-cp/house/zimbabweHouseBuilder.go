package house

type zimbabweHouseBuilder struct {
	house House
}

func (i *zimbabweHouseBuilder) buildWindow(numOfWindow int) {
	if numOfWindow > 2 {
		i.house.NumOfWindows = 2
	} else {
		i.house.NumOfWindows = numOfWindow
	}
}

func (i *zimbabweHouseBuilder) buildDoor() {
	i.house.NumOfDoors++
}

func (i *zimbabweHouseBuilder) buildGarage() {
	i.house.HasGarage = true
}

func (i *zimbabweHouseBuilder) buildSwimmingPool() {
	i.house.HasSwimmingPool = true
}

func (i *zimbabweHouseBuilder) getHouse() House {
	return i.house
}