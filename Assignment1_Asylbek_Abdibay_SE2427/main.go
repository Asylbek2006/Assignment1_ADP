package main

import (
	"fmt"
	"github.com/Asylbek2006/Assignment1/Employee"
	"github.com/Asylbek2006/Assignment1/Gym"
	"github.com/Asylbek2006/Assignment1/Hotel"
	"github.com/Asylbek2006/Assignment1/Wallet"
)

func main() {

	fmt.Println("--- Task 1: Hotel ---")
	hotel := Hotel.NewHotel()
	hotel.AddRoom(Hotel.Room{RoomNumber: 101, Type: "Single", PricePerNight: 50, IsOccupied: false})
	hotel.AddRoom(Hotel.Room{RoomNumber: 102, Type: "Double", PricePerNight: 100, IsOccupied: false})
	hotel.CheckIn(101)
	hotel.ListVacantRooms()

	fmt.Println("\n--- Task 2: Employee ---")
	employees := []Employee.Employee{
		Employee.FullTime{MonthlySalary: 3000, BonusRate: 0.1},
		Employee.PartTime{HourlyRate: 20, HoursWorked: 80},
		Employee.Intern{DailyRate: 50, DaysWorked: 20},
	}
	for _, e := range employees {
		fmt.Printf("Salary: $%.2f, Hours: %d\n", e.CalculateSalary(), e.GetWorkHours())
	}

	fmt.Println("\n--- Task 3: Gym ---")
	myGym := Gym.NewGym()
	myGym.AddMember(Gym.BasicMember{ID: 1, Name: "Ali"}, 1)
	myGym.AddMember(Gym.PremiumMember{ID: 2, Name: "Asel"}, 2)
	myGym.ListMembers()

	fmt.Println("\n--- Task 4: Wallet ---")
	myWallet := Wallet.Wallet{Owner: "Me", Balance: 0}
	myWallet.AddMoney(500)
	myWallet.SpendMoney(150)
	myWallet.GetBalance()
	myWallet.PrintHistory()
}
