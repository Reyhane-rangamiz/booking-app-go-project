package main

import (
	"fmt"
	"sync"
)


type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint 
}

var wg = sync.WaitGroup{}

func main() {


	greetUsers()

	for {


		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainedTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)  
			go sendTicket(userTickets, firstName, lastName, email)


			firstnamess := getFirstNames()
			fmt.Printf("bookings first names are: %v\n", firstnamess)

			if remainedTickets == 0 {
				fmt.Println("Sorry, all the tickets has been booked out :( \n Hope to have you with us next year")
				break
			}

		} else {
			fmt.Println("WARNING !!!")

			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("missing @gmail.com in your email address")
			}
			if !isValidTicketNumber {
				fmt.Println("the ticket count you entered doesn't match remaining ticket count or invalid number type")
			}
		}

	}

	wg.Wait() // blocks until the waitgroup counter is 0 
}

