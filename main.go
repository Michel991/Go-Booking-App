package main

import (
	"fmt"
	"time"
)

var conferenceName string = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct{
	firstName         string
	lastName          string
	email             string
	numberOfTickets   uint
}


func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			
			firstNames := getFirstName()
		    fmt.Printf("The First names of bookings are %v\n", firstNames)

		    if remainingTickets == 0 {
			    fmt.Println("Our conference is booked, come back Next Year")
			    break	
    	    }

        }  else {
		       if !isValidName {
			       fmt.Println("Firstname or lastname is too short")
		        }
			    if !isValidEmail {
			       fmt.Println("Email is incorrect, Please try again")
		        }
		        if !isValidTicketNumber {
			       fmt.Println("Ticket is above the available ticket, please try again")
		        }
	
		    }	

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have %v Tickets in total and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
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

	fmt.Println("Enter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&email)

	fmt.Println("Enter How Many Tickets You Want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n",bookings)
	
	fmt.Printf("Thank You %v %v For Booking %v Tickets, You Will Receive A Confirmation Email On Your Email %v\n", firstName, lastName, userTickets, email)
	
	fmt.Printf("We Have %v Remaining Tickets For This %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets was sent to %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending Ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("#####################")
}

