package main

import (
	"fmt"
	"strings"
	"BookingApp/helper"
)

//Package variables
var conferenceName string = "Great Inc"
const conferenceTickets uint8 = 50
var remainingTickets uint8 = 50
var bookings = []string{}



func main() {
	//variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint8
	

	for remainingTickets != 0 {

		helper.GreetUsers(conferenceName,conferenceTickets,remainingTickets)

		firstName, lastName = helper.GetNameInput()

		var validFullName = helper.CheckNameInput(firstName,lastName)

		email = helper.GetMailInput()

		var validMail = helper.CheckMailInput(email)

		userTickets = helper.GetTicketInput()

		var validTickets = helper.CheckTicketInput(userTickets, remainingTickets)

		remainingTickets = remainingTickets - validTickets

		bookings = append(bookings, validFullName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", validFullName, validTickets, validMail)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Println("============================================")
		fmt.Printf("The first names of bookings are: %v \n", getFirstNames(bookings))
		fmt.Println("============================================")

		//bookTicket(remainingTickets,validTickets,bookings,validFullName,validMail,conferenceName)
	}

	helper.ShowEndMessage()
}

//funcoes

func getFirstNames(bookings []string) []string {
	var firstNames = []string{}

	for _, nomeDoInscrito := range bookings {
		var primeiroNome = strings.Fields(nomeDoInscrito)
		firstNames = append(firstNames, primeiroNome[0])
	}

	return firstNames
}





/* func bookTicket(_remainingTickets uint8,validTickets uint8,bookings []string,validFullName string,validMail string,conferenceName string) {
	remainingTickets = _remainingTickets - validTickets

	bookings = append(bookings, validFullName)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", validFullName, validTickets, validMail)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Println("============================================")
	fmt.Printf("The first names of bookings are: %v \n", getFirstNames(bookings))
	fmt.Println("============================================")
} */

