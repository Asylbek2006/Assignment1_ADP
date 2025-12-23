package Hotel

import "fmt"

type Room struct {
	RoomNumber    int
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Rooms map[int]Room
}

func NewHotel() *Hotel {
	return &Hotel{Rooms: make(map[int]Room)}
}

func (h *Hotel) AddRoom(r Room) {
	h.Rooms[r.RoomNumber] = r
}

func (h *Hotel) CheckIn(roomNum int) {
	if room, ok := h.Rooms[roomNum]; ok {
		if !room.IsOccupied {
			room.IsOccupied = true
			h.Rooms[roomNum] = room
			fmt.Printf("Room %d checked in successfully.\n", roomNum)
		} else {
			fmt.Println("Room is already occupied.")
		}
	} else {
		fmt.Println("Room not found.")
	}
}

func (h *Hotel) CheckOut(roomNum int) {
	if room, ok := h.Rooms[roomNum]; ok {
		room.IsOccupied = false
		h.Rooms[roomNum] = room
		fmt.Printf("Room %d checked out.\n", roomNum)
	}
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant Rooms:")
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Printf("Room %d (%s) - $%.2f\n", room.RoomNumber, room.Type, room.PricePerNight)
		}
	}
}
