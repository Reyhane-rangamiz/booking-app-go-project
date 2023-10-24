package main

import (
	"fmt"
	"strings"
	"time"

) 


const conferanceTickets int = 50
var confranceName = "Debi's conferance"
var remainedTickets uint = 50 
var bookings = make([]UserData, 0)  // initiates a userdata struct

  
func greetUsers() {

	fmt.Printf("welcome to %v please fill the form below\n", confranceName)
	fmt.Printf("%v tickets are left from a total of %v \n", remainedTickets, conferanceTickets)
	fmt.Println("get your tickets here to attend")
}


func validateUserInput(fName string, lName string, email string, userTickets uint, remainedTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainedTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please Enter Your Name:")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter Your Lastname:")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter Your email:")
	fmt.Scan(&email)

	fmt.Println("Tickets Count:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainedTickets = remainedTickets - userTickets


var UserData = UserData{
	firstName : firstName,
	lastName : lastName,
	email :  email,
	numberOfTickets: userTickets,
}

	bookings = append(bookings, UserData)

	fmt.Println("there are", remainedTickets, "left")
	fmt.Printf("congerats, You Just Booked %v tickets from %v, %v more are remained\n", userTickets, confranceName, remainedTickets)
	fmt.Printf("the whole array is : %v\n", bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	// sprintf is a way to print sth and save it in a variable 
	fmt.Println("#######################")
	fmt.Printf("sending ticket...  \n Sending the verification of %v to %v email address \n", ticket, email)
	fmt.Println("#######################")

	wg.Done() 
}