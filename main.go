package main

import (
	"BookingApp/helper"
	"fmt"
	"time"
	"sync"
)

//Package variables
var conferenceName string = "Great Inc"

const conferenceTickets uint8 = 50

var remainingTickets uint8 = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

var bookings = make([]UserData, 0)

var counter uint8 = 0
var wg = sync.WaitGroup{}
func main() {
	//variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint8
	var firstNamesList []string = []string{}
	for remainingTickets != 0 {

		helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

		firstName, lastName = helper.GetNameInput()

		var validFullName = helper.CheckNameInput(firstName, lastName)

		email = helper.GetMailInput()

		var validMail = helper.CheckMailInput(email)

		userTickets = helper.GetTicketInput()

		var validTickets = helper.CheckTicketInput(userTickets, remainingTickets)

		remainingTickets = remainingTickets - validTickets

		var userData = UserData{
			firstName:       helper.SplitName(validFullName)[0],
			lastName:        helper.SplitName(validFullName)[1],
			email:           validMail,
			numberOfTickets: validTickets,
		}

		bookings = append(bookings, userData)

		firstNamesList = append(firstNamesList, bookings[counter].firstName)
		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", validFullName, validTickets, validMail)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Println("============================================")
		fmt.Println(firstNamesList)
		fmt.Println("============================================")
		wg.Add(1)
		go sendTicket(userData.numberOfTickets,userData.firstName,userData.lastName,userData.email)

		counter++
	}

	helper.ShowEndMessage()
	wg.Wait()
}

func sendTicket(userTickets uint8,firstName string, lastName string,email string ){
	time.Sleep(40*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)
	fmt.Println("==============")
	fmt.Printf("Sending ticket:\n %v \n to email addres %v\n",ticket,email)
	fmt.Println("==============")
	wg.Done()
}