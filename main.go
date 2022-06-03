package main

import (
	"fmt"
)

func main(){
	// variable
	var conferenceName ="Go Booking"
	// constant
	const conferenceTicket=50;

	var remainingTickets=50;

	fmt.Printf("Welcome to %v app\n", conferenceName);
	fmt.Println("We have total of",conferenceTicket, "tickets and",remainingTickets,"are remaining")
	fmt.Println("Get your tickets here to attend");

	// data types in go
	var firstName string
	var lastName string
	var email string
	var userTicket int


	// Arrays- cannot mix the types in array
	// var bookings = [50]string{}
	var bookings [50]string



	// type of
	// fmt.Printf("conference name is %T\n",conferenceName)

	// reading user input
	fmt.Println("Enter you first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTicket)

	remainingTickets=remainingTickets-userTicket

    // assigning value to array
	bookings[0]=firstName + " " + lastName

	fmt.Printf("the whole array: %v\n",bookings)
	fmt.Printf("the first value: %v\n",bookings[0])
	fmt.Printf("the length is: %v\n",len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets .Email is send to %v\n",firstName,lastName,userTicket,email)
	fmt.Printf("%v tickets are remaining now", remainingTickets)


}
