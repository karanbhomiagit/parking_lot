package parkingLot

import (
	"testing"

	"github.com/karanbhomiagit/parking_lot/src/slot"
	"github.com/karanbhomiagit/parking_lot/src/vehicle"
	"github.com/stretchr/testify/assert"
)

func TestNewParkingLot(t *testing.T) {

	t.Run("Successfully create parking lot for size > 1", func(t *testing.T) {
		p := NewParkingLot(2)
		assert := assert.New(t)
		assert.Equal(len(p), 2)
		assert.Equal(p[0].IsFree(), true)
		assert.Equal(p[1].IsFree(), true)
	})

	t.Run("Return nil when trying to create parking lot for size 0", func(t *testing.T) {
		p := NewParkingLot(0)
		assert := assert.New(t)
		assert.Nil(p)
	})
}

func TestPark(t *testing.T) {

	t.Run("Successfully park vehicles and return error message when capacity is full", func(t *testing.T) {
		p := NewParkingLot(2)
		assert := assert.New(t)
		v := new(vehicle.Vehicle)
		res, err := p.Park(*v)
		assert.Nil(err)
		assert.Equal(res, 1)
		res, err = p.Park(*v)
		assert.Nil(err)
		assert.Equal(res, 2)
		res, err = p.Park(*v)
		assert.Equal(res, -1)
		assert.Equal(err.Error(), "No free slots found")
	})

	t.Run("Successfully park in first empty slot", func(t *testing.T) {
		p := NewParkingLot(3)
		assert := assert.New(t)

		v := new(vehicle.Vehicle)
		(*v).SetRegistrationNumber("abcd")
		(*v).SetColor("Red")
		res, err := p.Park(*v)
		assert.Nil(err)
		assert.Equal(res, 1)

		v1 := new(vehicle.Vehicle)
		(*v1).SetRegistrationNumber("pqrs")
		(*v1).SetColor("Black")
		res, err = p.Park(*v1)
		assert.Equal(res, 2)
		assert.Nil(err)

		err = p.Leave(1)
		assert.Nil(err)

		v2 := new(vehicle.Vehicle)
		(*v2).SetRegistrationNumber("djoj")
		(*v2).SetColor("Green")
		res, err = p.Park(*v2)
		assert.Equal(res, 1)
		assert.Nil(err)
	})

}

func TestLeave(t *testing.T) {

	t.Run("Return error if try to leave a slot number not present", func(t *testing.T) {
		p := NewParkingLot(2)
		assert := assert.New(t)
		err := p.Leave(3)
		assert.Equal(err.Error(), "Enter slotNumber less than length of parking lot")
	})

	t.Run("Return error if try to leave a slot number < 1", func(t *testing.T) {
		p := NewParkingLot(2)
		assert := assert.New(t)
		err := p.Leave(0)
		assert.Equal(err.Error(), "Enter slotNumber > 0")
	})
}

func TestGetStatus(t *testing.T) {

	t.Run("Successfully get status of empty parking slot", func(t *testing.T) {
		p := NewParkingLot(2)
		assert := assert.New(t)
		occupiedSlots := p.GetStatus()
		assert.Equal(occupiedSlots, []slot.Slot{})
	})

	t.Run("Successfully get status of randomly filled parking slot", func(t *testing.T) {
		p := NewParkingLot(4)
		assert := assert.New(t)

		v := new(vehicle.Vehicle)
		(*v).SetRegistrationNumber("abcd")
		(*v).SetColor("Red")

		res, err := p.Park(*v)
		assert.Nil(err)
		assert.Equal(res, 1)

		v1 := new(vehicle.Vehicle)
		(*v1).SetRegistrationNumber("pqrs")
		(*v1).SetColor("Black")
		res, err = p.Park(*v1)
		assert.Equal(res, 2)
		assert.Nil(err)

		v2 := new(vehicle.Vehicle)
		(*v2).SetRegistrationNumber("djoj")
		(*v2).SetColor("Green")
		res, err = p.Park(*v2)
		assert.Equal(res, 3)
		assert.Nil(err)

		err = p.Leave(2)
		assert.Nil(err)
		err = p.Leave(2)
		assert.Nil(err)

		occupiedSlots := p.GetStatus()
		assert.Equal(len(occupiedSlots), 2)
	})

}

func TestGetSlotNumbersByColor(t *testing.T) {

	t.Run("Successfully get slot numbers for multiple cars with same color", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		nums := p.GetSlotNumbersByColor("Green")
		assert.Equal(nums, []int{3, 4})
	})

	t.Run("Successfully get slot numbers for 1 car with a color", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		nums := p.GetSlotNumbersByColor("Black")
		assert.Equal(nums, []int{6})
	})

	t.Run("Successfully get slot numbers for an absent color", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		nums := p.GetSlotNumbersByColor("Magenta")
		assert.Equal(nums, []int{})
	})

}

func TestGetRegNumbersByColor(t *testing.T) {

	t.Run("Successfully get reg numbers for multiple cars with same color", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		nums := p.GetRegNumbersByColor("Green")
		assert.Equal(nums, []string{"djoj", "dokdoeke"})
	})

	t.Run("Successfully get reg numbers for 1 car with a color", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		nums := p.GetRegNumbersByColor("Black")
		assert.Equal(nums, []string{"929292"})
	})

	t.Run("Successfully get reg numbers for an absent color", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		nums := p.GetRegNumbersByColor("Magenta")
		assert.Equal(nums, []string{})
	})

}

func TestGetSlotNumberByRegNum(t *testing.T) {

	t.Run("Successfully get slot number for a reg num present", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		num, err := p.GetSlotNumberByRegNum("929292")
		assert.Nil(err)
		assert.Equal(num, 5)
	})

	t.Run("Successfully get error for an absent reg num", func(t *testing.T) {
		p := GetRandomlyFilledParkingLot()
		assert := assert.New(t)
		_, err := p.GetSlotNumberByRegNum("xxx")
		assert.Equal(err.Error(), "Not found")
	})

}

func GetRandomlyFilledParkingLot() ParkingLot {
	p := NewParkingLot(6)

	v := &vehicle.Vehicle{
		Color:              "Red",
		RegistrationNumber: "abcd",
	}
	p.Park(*v)

	v = &vehicle.Vehicle{
		Color:              "Black",
		RegistrationNumber: "pqrs",
	}
	p.Park(*v)

	v = &vehicle.Vehicle{
		Color:              "Green",
		RegistrationNumber: "djoj",
	}
	p.Park(*v)

	v = &vehicle.Vehicle{
		Color:              "Green",
		RegistrationNumber: "dokdoeke",
	}
	p.Park(*v)

	v = &vehicle.Vehicle{
		Color:              "Red",
		RegistrationNumber: "828282",
	}
	p.Park(*v)

	v = &vehicle.Vehicle{
		Color:              "Black",
		RegistrationNumber: "929292",
	}
	p.Park(*v)

	p.Leave(2)

	return p
}
