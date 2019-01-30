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
