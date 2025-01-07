package main

import "fmt"

func main() {
	// variables
	var userName string
	var ticketsLeft uint
	var userTickets uint
	var bookings []string
	ticketsLeft = 18

	fmt.Println("Who uses a cli Car booking app?")
	fmt.Println("                               ")

	for {
		fmt.Println("Enter nickname: ")
		fmt.Scan(&userName)
		bookings = append(bookings, userName)
		fmt.Printf("Welcome onboard %v!\n", userName)
		fmt.Printf("So %v we have %v tickets left\n", userName, ticketsLeft)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		fmt.Printf("%v you just booked %v tickets!\n", userName, userTickets)

		ticketsLeft = ticketsLeft - userTickets
		fmt.Println("users: ", bookings)
		fmt.Printf("%v tickets remaining to be booked!\n", ticketsLeft)

		if ticketsLeft == 0 {
			break
		}
	}
	// user provided values.

}
