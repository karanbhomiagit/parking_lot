package slot

import "github.com/karanbhomiagit/parking_lot/src/vehicle"

type Slot struct {
	vehicle vehicle.Vehicle
	number  int
	free    bool
}

func (s *Slot) isFree() bool {
	return (*s).free
}

func (s *Slot) setFree(isFree bool) {
	(*s).free = isFree
}

func (s *Slot) getNumber() int {
	return (*s).number
}

func (s *Slot) getVehicle() vehicle.Vehicle {
	return (*s).vehicle
}

func (s *Slot) setVehicle(v vehicle.Vehicle) {
	(*s).vehicle = v
}
