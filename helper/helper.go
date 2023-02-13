package helper

import (
	"fmt"
	"strings"
)

func GetNameInput() (string, string) {
	var firstName string
	var lastName string
	fmt.Println("Enter your first name: ")

	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")

	fmt.Scan(&lastName)

	return firstName, lastName
}
func ShowEndMessage() {
	fmt.Println("Booking Closed, no tickets left")
	fmt.Println("================End of Application================")
}
func GreetUsers(conferenceName string, conferenceTickets uint8, remainingTickets uint8) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining \n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")
}
func GetMailInput() string {
	var email string

	fmt.Println("Enter your email address: ")

	fmt.Scan(&email)

	return email
}
func GetTicketInput() uint8 {
	var userTickets uint8
	fmt.Println("Enter Number of tickets: ")
	fmt.Scan(&userTickets)

	return userTickets

}
func CheckNameInput(firstName string, lastName string) string {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	for !isValidName {
		fmt.Println("enter a valid first name: (2 characters or more)")
		fmt.Scan(&firstName)
		fmt.Println("enter a valid last name: (2 characters or more)")
		fmt.Scan(&lastName)
		isValidName = len(firstName) >= 2 && len(lastName) >= 2
	}
	return firstName + " " + lastName

}

func CheckMailInput(email string) string {
	var isValidMail = strings.Contains(email, "@")
	for !isValidMail {
		fmt.Println("enter a valid email address: (must contain '@')")
		fmt.Scan(&email)
		isValidMail = strings.Contains(email, "@")
	}
	return email
}

func CheckTicketInput(userTickets uint8, avaiableTickets uint8) uint8 {

	for userTickets <= 0 || userTickets > avaiableTickets {
		fmt.Printf("There are only %v remaining and you cannot buy 0 tickets! type again: ", avaiableTickets)
		fmt.Scan(&userTickets)
	}

	return userTickets

}
