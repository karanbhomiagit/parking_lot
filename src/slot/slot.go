package slot

import "github.com/karanbhomiagit/parking_lot/src/vehicle"

type Slot struct {
	vehicle vehicle.Vehicle
	number  int
	free    bool
}

func (s *Slot) IsFree() bool {
	return (*s).free
}

func (s *Slot) SetFree(isFree bool) {
	(*s).free = isFree
}

func (s *Slot) SetNumber(num int) {
	(*s).number = num
}

func (s *Slot) GetNumber() int {
	return (*s).number
}

func (s *Slot) GetVehicle() vehicle.Vehicle {
	return (*s).vehicle
}

func (s *Slot) SetVehicle(v vehicle.Vehicle) {
	(*s).vehicle = v
}
