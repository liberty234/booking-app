package main

import "fmt"

func main() {
	var comferenceName = "Big Doris App"
	const comferenceTickets = 50
	var RemainingTickets uint = 50

	fmt.Printf("Welcome to %v\n", comferenceName)
	fmt.Printf("we have a total tickets of %v and %v are still Available.\n", comferenceTickets, RemainingTickets)
	fmt.Println("Get your tickeks to Attend.")

	var firstName string
	var lastName string
	var email string
	var phoneNumber int
	var age int
	var tickets uint

	RemainingTickets = RemainingTickets - tickets

	//asking userName
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last  name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your phonme numbers:")
	fmt.Scan(&phoneNumber)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	fmt.Println("Enter Tickets:")
	fmt.Scan(&tickets)

	fmt.Printf("thank you %v %v  for booking %v tickets with us,you will recieve a comfirmation email at %v\n", firstName, lastName, tickets, email)
	fmt.Printf("we have %v remaining", RemainingTickets-tickets)

}
