package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/karanbhomiagit/parking_lot/src/parkingLot"
	"github.com/karanbhomiagit/parking_lot/src/vehicle"
)

var p parkingLot.ParkingLot

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) > 1 {
		//log.Println("  ", os.Args)
		file, err := os.Open(os.Args[1])
		if err != nil {
			//log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		switch command[0] {
		case "create_parking_lot":
			if len(command) < 2 {
				log.Println("Missing arguments for create_parking_lot. Format: create_parking_lot <length>")
			} else {
				length, err := strconv.Atoi(command[1])
				if err != nil {
					log.Println("Please enter a number as length.")
				} else {
					p = parkingLot.NewParkingLot(length)
					fmt.Fprintf(os.Stdout, "Created a parking lot with %d slots. \n", length)
				}
			}

		case "park":
			if len(command) < 3 {
				log.Println("Missing arguments for park. Format: park <registrationNumber> <color>")
			} else if p == nil {
				log.Println("Please create the parking lot as first step.")
			} else {
				slotNumber, err := p.Park(vehicle.Vehicle{
					RegistrationNumber: command[1],
					Color:              command[2],
				})
				if err != nil {
					fmt.Fprintln(os.Stdout, err.Error())
				} else {
					fmt.Fprintln(os.Stdout, "Allocated slot number:", slotNumber)
				}
			}

		case "status":
			if p == nil {
				log.Println("Please create the parking lot as first step.")
			} else {
				slots := p.GetStatus()
				fmt.Fprintf(os.Stdout, "Slot No.    Registration No    Colour\n")
				for _, slot := range slots {
					fmt.Fprintf(os.Stdout, "%d           %s      %s\n", slot.GetNumber(), slot.GetVehicle().GetRegistrationNumber(), slot.GetVehicle().GetColor())
				}
			}

		case "leave":
			if len(command) < 2 {
				log.Println("Missing arguments for leave. Format: leave <slotNumber>")
			} else if p == nil {
				log.Println("Please create the parking lot as first step.")
			} else {
				slotId, err := strconv.Atoi(command[1])
				if err != nil {
					log.Println("Please enter a number as slotNumber.")
				} else {
					err := p.Leave(slotId)
					if err != nil {
						fmt.Fprintln(os.Stdout, err.Error())
					} else {
						fmt.Fprintf(os.Stdout, "Slot number %d is free \n", slotId)
					}
				}
			}

		case "registration_numbers_for_cars_with_colour":
			if len(command) < 2 {
				log.Println("Missing arguments for registration_numbers_for_cars_with_colour. Format: registration_numbers_for_cars_with_colour <color>")
			} else if p == nil {
				log.Println("Please create the parking lot as first step.")
			} else {
				regNums := p.GetRegNumbersByColor(command[1])
				fmt.Fprintln(os.Stdout, strings.Join(regNums, ", "))
			}

		case "slot_numbers_for_cars_with_colour":
			if len(command) < 2 {
				log.Println("Missing arguments for slot_numbers_for_cars_with_colour. Format: slot_numbers_for_cars_with_colour <color>")
			} else if p == nil {
				log.Println("Please create the parking lot as first step.")
			} else {
				slots := p.GetSlotNumbersByColor(command[1])
				fmt.Fprintln(os.Stdout, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slots)), ", "), "[]"))
			}

		case "slot_number_for_registration_number":
			if len(command) < 2 {
				log.Println("Missing arguments for slot_number_for_registration_number. Format: slot_number_for_registration_number <registrationNumber>")
			} else if p == nil {
				log.Println("Please create the parking lot as first step.")
			} else {
				slotNum, err := p.GetSlotNumberByRegNum(command[1])
				if err != nil {
					fmt.Fprintln(os.Stdout, err.Error())
				} else {
					fmt.Fprintln(os.Stdout, slotNum)
				}
			}

		default:
			fmt.Fprintln(os.Stdout, "Unknown Command")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
