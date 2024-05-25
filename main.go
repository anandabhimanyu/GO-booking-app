package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTicket = 50

var conferenceName = "Go Confefence"
var reamainingTicket uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := gteUserInput()
	isValidNmae, isValidEmail, isValidTicketNumber := validUserInput(firstName, lastName, email, userTickets, reamainingTicket)
	//Check for User cant not buy more then Remaining Tickets.
	if isValidNmae && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		//Add() Sets the number  of goroutines to wait  ..
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		// Call Funtion to Print First Name
		firstNames := getFirstNames()
		fmt.Printf("The First Names of  bookings are :%v\n", firstNames)

		//Check for is Remaning tickets is Zero or not
		if reamainingTicket == 0 {
			//end the Program.
			fmt.Println("Our Conference is booked out. Come back next year.")
			//break
		}
	} else {
		if !isValidNmae {
			fmt.Println("First name or last name you  entered is too short ")
		}
		if !isValidEmail {
			fmt.Println("Emaid address should contaion @")
		}
		if !isValidNmae {
			fmt.Println("Number of Tickets you entered is invalid")
		}
	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to  %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v Tickets And  %v are still available.\n", conferenceTicket, reamainingTicket)
	fmt.Println("Gte your Tickets here to Attend")

}

func getFirstNames() []string {
	//Iterates Over Elements of array Used Range Keyword..
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func gteUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//Ask the user for their Name
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)

	fmt.Println("Enter nu of  Tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	reamainingTicket = reamainingTicket - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("Thank you %v %v for bookings %v tickets. You will receive a conformationemailat %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v\n", reamainingTicket, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending ticket:\n %v \nto the email address %v\n", ticket, email)
	fmt.Println("###########################")
	wg.Done()

}
