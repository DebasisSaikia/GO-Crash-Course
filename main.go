package main

import "fmt"

func main(){
	// variable
	var conferenceName ="Go Booking"
	// constant
	const conferenceTicket=50;

	var remainingTickets=50;

	fmt.Printf("Welcome to %v app\n", conferenceName);
	fmt.Println("We have total of",conferenceTicket, "tickets and",remainingTickets,"are remaining")
	fmt.Println("Get your tickets here to attend");

}
