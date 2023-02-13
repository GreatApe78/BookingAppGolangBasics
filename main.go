package main

import (
	"BookingApp/helper"
	"fmt"
	"strconv"
	
)

//Package variables
var conferenceName string = "Great Inc"
const conferenceTickets uint8 = 50
var remainingTickets uint8 = 50
var bookings = make([]map[string]string,0)

var counter uint8 = 0 


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

		//create a map for user

		var userData = make(map[string]string)
		userData["firstName"] = helper.SplitName(validFullName)[0] 
		userData["lastName"] = helper.SplitName(validFullName)[1]
		userData["email"] = validMail
		userData["numberOfTickets"] = strconv.FormatUint(uint64(validTickets),10)

		bookings = append(bookings, userData)
		
		
		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", validFullName, validTickets, validMail)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Println("============================================")
		fmt.Println(bookings)
		fmt.Println("============================================")

		counter++
	}

	helper.ShowEndMessage()
}

//funcoes




