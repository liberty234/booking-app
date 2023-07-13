package main

import (
	"fmt"
)

var comferenceName = "Big Doris App"

const comferenceTickets = 50

var RemainingTickets uint = 50

var firstName string
var lastName string
var email string
var phoneNumber string
var userAge int
var tickets uint

func getFirstName() {

	//var firstName string
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

}

func getLastName() {

	//var lastName string
	fmt.Println("Enter your last  name:")
	fmt.Scan(&lastName)

}
func getEmail() {

	//var email string
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

}

func getPhoneNumber() {

	//var phoneNumber int

	fmt.Println("Enter your phone number:")
	fmt.Scan(&phoneNumber)

}

func getAge() {

	//var userAge int

	fmt.Println("Enter your age:")
	fmt.Scan(&userAge)

}

func getTickets() {

	//var tickets uint

	fmt.Println("Enter Tickets:")
	fmt.Scan(&tickets)
	RemainingTickets = RemainingTickets - tickets

}

func main() {

	fmt.Printf("Welcome to %v\n", comferenceName)
	fmt.Printf("we have a total tickets of %v and %v are still Available.\n", comferenceTickets, RemainingTickets)
	fmt.Println("Get your tickeks to Attend.")

	//ask User thier information
	getFirstName()
	getLastName()
	getPhoneNumber()
	getEmail()
	getAge()
	getTickets()

	fmt.Printf("Thank you %s %s for booking %v tickets with us, you will recieve a comfirmation email at %s\n ", firstName, lastName, tickets, email)
	fmt.Printf("we have %v remaining\n", RemainingTickets)

}
