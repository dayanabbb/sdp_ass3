package main

import (
	"fmt"
	"sync"
)

type HotelManager struct {
	Rooms map[int]string
}

var instance *HotelManager
var once sync.Once

func GetHotelManagerInstance() *HotelManager {
	once.Do(func() {
		instance = &HotelManager{
			Rooms: make(map[int]string),
		}
	})
	return instance
}

func (hm *HotelManager) CheckIn(roomNumber int, guestName string) {
	if occupant, ok := hm.Rooms[roomNumber]; !ok {
		hm.Rooms[roomNumber] = guestName
		fmt.Printf("%s has checked into room %d.\n", guestName, roomNumber)
	} else {
		fmt.Printf("room %d is already occupied by %s.\n", roomNumber, occupant)
	}
}

func (hm *HotelManager) CheckOut(roomNumber int) {
	if occupant, ok := hm.Rooms[roomNumber]; ok {
		delete(hm.Rooms, roomNumber)
		fmt.Printf("%s has checked out of room %d.\n", occupant, roomNumber)
	} else {
		fmt.Printf("room %d is currently unoccupied.\n", roomNumber)
	}
}

func main() {
	hotel1 := GetHotelManagerInstance()
	hotel1.CheckIn(101, "Raushan")

	hotel2 := GetHotelManagerInstance()
	hotel2.CheckIn(101, "Dayana")

	hotel1.CheckOut(101)
	hotel2.CheckOut(102)
}
