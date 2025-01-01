package main

import "fmt"

func main()  {
	// variables
	var userName string
	// var totalTickets = 18/
	var ticketsLeft:uint = 18
	var userTickets uint

	fmt.Println("Who uses a cli Car booking app?")
	fmt.Println("                                 ")
	// user variables
	userName = "favour"
	userTickets = 2
	fmt.Printf("Welcome onboard %v!\n", userName)
	fmt.Printf("So %v we have %v tickets left\n", userName, ticketsLeft)
	fmt.Printf("%v you just booked %v tickets!\n", userName, userTickets)
	// ticketsLeft = ticketsLeft - userTickets

}

// func Ticketleft(totalTickets, ticketsLeft, userTickets){
// 	totalTickets = ticketsLeft - userTickets
// 	ticketsLeft = totalTickets

// 	return ticketsLeft
// }