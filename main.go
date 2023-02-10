package main

import (
	"fmt"
	"strings"
)

func main() {
	//variables
	var conferenceName string = "Great Inc"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint8
	var bookings = []string{}
	var pf = fmt.Printf
	var pl = fmt.Println

	for remainingTickets != 0 {

		greetUsers(conferenceName, conferenceTickets, remainingTickets)

		pf("We have a total of %v tickets and %v tickets remaining \n", conferenceTickets, remainingTickets)

		pl("Get your tickets here to attend")

		firstName,lastName = getNameInput()

		var validFullName = checkNameInput(firstName, lastName)

		email = getMailInput()

		var validMail = checkMailInput(email)

		userTickets = getTicketInput()


		var validTickets = checkTicketInput(userTickets, remainingTickets)

		remainingTickets = remainingTickets - validTickets

		bookings = append(bookings, validFullName)

		pf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", validFullName, validTickets, validMail)

		pf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Println("============================================")
		fmt.Printf("The first names of bookings are: %v \n", getFirstNames(bookings))
		fmt.Println("============================================")
	}

	pl("Booking Closed, no tickets left")
	fmt.Println("================End of Application================")
}

//funcoes
func greetUsers(conferenceName string, conferenceTickets uint8, remainingTickets uint8) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining \n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	var firstNames = []string{}

	for _, nomeDoInscrito := range bookings {
		var primeiroNome = strings.Fields(nomeDoInscrito)
		firstNames = append(firstNames, primeiroNome[0])
	}

	return firstNames
}

func checkNameInput(firstName string, lastName string) string {
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

func checkMailInput(email string) string {
	var isValidMail = strings.Contains(email, "@")
	for !isValidMail {
		fmt.Println("enter a valid email address: (must contain '@')")
		fmt.Scan(&email)
		isValidMail = strings.Contains(email, "@")
	}
	return email
}

func checkTicketInput(userTickets uint8, avaiableTickets uint8) uint8 {
	var isValidTickets = userTickets > 0
	for !isValidTickets {
		fmt.Println("You need to buy at least one ticket! type again:")
		fmt.Scan(&userTickets)
		isValidTickets = userTickets > 0
	}
	for userTickets > avaiableTickets {
		fmt.Printf("You cannot by more than the %v avaiable tickets! type again:\n", avaiableTickets)
		fmt.Scan(&userTickets)
	}
	return userTickets
}

func getNameInput() (string, string) {
	var firstName string
	var lastName string
	fmt.Println("Enter your first name: ")

	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")

	fmt.Scan(&lastName)

	return firstName,lastName
}

func getMailInput() string{
	var email string
	
	fmt.Println("Enter your email address: ")
	
	fmt.Scan(&email)

	return email
}
func getTicketInput() uint8{
	var userTickets uint8
	fmt.Println("Enter Number of tickets: ")
	fmt.Scan(&userTickets)

	return userTickets

}