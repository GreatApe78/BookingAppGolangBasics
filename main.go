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
	var input = fmt.Scan

	for remainingTickets != 0 {

		greetUsers(conferenceName, conferenceTickets, remainingTickets)

		pf("We have a total of %v tickets and %v tickets remaining \n", conferenceTickets, remainingTickets)

		pl("Get your tickets here to attend")

		pl("Enter your first name: ")

		input(&firstName)

		pl("Enter your last name: ")

		input(&lastName)

		var isValidName = len(firstName) >= 2 && len(lastName) >= 2

		for !isValidName {
			pl("enter a valid first name: (2 characters or more)")
			input(&firstName)
			pl("enter a valid last name: (2 characters or more)")
			input(&lastName)
			isValidName = len(firstName) >= 2 && len(lastName) >= 2
		}

		pl("Enter your email address: ")

		input(&email)

		var isValidMail = strings.Contains(email, "@")
		for !isValidMail {
			pl("enter a valid email address: (must contain '@')")
			input(&email)
			isValidMail = strings.Contains(email, "@")
		}

		pl("How many tickets do you want to buy?")

		input(&userTickets)

		var isValidTickets = userTickets > 0

		for !isValidTickets {
			pl("You need to buy at least one ticket!")
			input(&userTickets)
			isValidTickets = userTickets > 0
		}

		for userTickets > remainingTickets {
			pf("You cannot buy more tickets than the remaining ones. Choose a valid quantity! %v or less\n", remainingTickets)
			fmt.Scan(&userTickets)
		}

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		pf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		pf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("The first names of bookings are: %v \n",getFirstNames(bookings))
	}

	pl("Booking Closed, no tickets left")
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

func validateNameInput(firstName string,lastName string) bool{
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		for !isValidName {
			fmt.Println("enter a valid first name: (2 characters or more)")
			fmt.Scan(&firstName)
			fmt.Println("enter a valid last name: (2 characters or more)")
			fmt.Scan(&lastName)
			isValidName = len(firstName) >= 2 && len(lastName) >= 2
		}
	return isValidName

}