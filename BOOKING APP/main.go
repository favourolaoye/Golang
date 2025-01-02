package main

import "fmt"

func main()  {
	// variables
	var userName string
	var ticketsLeft uint
	var userTickets uint

	ticketsLeft = 18

	fmt.Println("Who uses a cli Car booking app?")
	fmt.Println("                               ")
	// user provided values.
	fmt.Println("Enter nickname: ");
	fmt.Scan(&userName);

	fmt.Printf("Welcome onboard %v!\n", userName)
	fmt.Printf("So %v we have %v tickets left\n", userName, ticketsLeft)

	fmt.Println("Enter number of tickets: ");
	fmt.Scan(&userTickets);
	
	fmt.Printf("%v you just booked %v tickets!\n", userName, userTickets)
	// ticketsLeft = ticketsLeft - userTickets

}
  
// func Ticketleft(totalTickets, ticketsLeft, userTickets){
// 	totalTickets = ticketsLeft - userTickets
// 	ticketsLeft = totalTickets

// 	return ticketsLeft
// } n kjllkk