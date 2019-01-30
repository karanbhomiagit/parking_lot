package parkingLot

import (
	"errors"

	"github.com/karanbhomiagit/parking_lot/src/slot"
	"github.com/karanbhomiagit/parking_lot/src/vehicle"
)

type ParkingLot []slot.Slot

func NewParkingLot(numberOfSlots int) ParkingLot {
	if numberOfSlots < 1 {
		return nil
	}
	slots := make([]slot.Slot, numberOfSlots)
	for i, slot := range slots {
		(&slot).SetNumber(i)
		(&slot).SetFree(true)
		slots[i] = slot
	}
	var p ParkingLot = slots
	return p
}

func (p *ParkingLot) GetStatus() []slot.Slot {
	occupiedSlots := []slot.Slot{}
	for _, slot := range *p {
		if (&slot).IsFree() == false {
			occupiedSlots = append(occupiedSlots, slot)
		}
	}
	// fmt.Println("Status", occupiedSlots)
	return occupiedSlots
}

func (p *ParkingLot) Park(v vehicle.Vehicle) (int, error) {
	// fmt.Printf("park %+v", car)
	for i, slot := range *p {
		if (&slot).IsFree() == true {
			(&slot).SetVehicle(v)
			(&slot).SetFree(false)
			(*p)[i] = slot
			return i + 1, nil
		}
	}
	return -1, errors.New("Sorry, parking lot is full")
}

func (p *ParkingLot) Leave(slotNum int) error {
	if slotNum < 1 {
		return errors.New("Enter slotNumber > 0")
	}
	if slotNum > len(*p) {
		return errors.New("Enter slotNumber less than length of parking lot")
	}
	(*p)[slotNum-1].SetFree(true)
	return nil
}

func (p *ParkingLot) GetSlotNumbersByColor(color string) []int {
	slotNumbers := []int{}
	for _, slot := range *p {
		if slot.IsFree() == false && slot.GetVehicle().GetColor() == color {
			slotNumbers = append(slotNumbers, slot.GetNumber()+1)
		}
	}
	return slotNumbers
}

func (p *ParkingLot) GetRegNumbersByColor(color string) []string {
	regNumbers := []string{}
	for _, slot := range *p {
		if slot.IsFree() == false && slot.GetVehicle().GetColor() == color {
			regNumbers = append(regNumbers, slot.GetVehicle().GetRegistrationNumber())
		}
	}
	return regNumbers
}

func (p *ParkingLot) GetSlotNumberByRegNum(regNum string) (int, error) {
	for _, slot := range *p {
		if slot.IsFree() == false && slot.GetVehicle().GetRegistrationNumber() == regNum {
			return slot.GetNumber(), nil
		}
	}
	return -1, errors.New("Not found")
}
