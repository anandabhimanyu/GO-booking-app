package main

import "strings"

// reamainingTicket
func validUserInput(firstName string, lastName string, email string, userTickets uint, reamainingTicket uint) (bool, bool, bool) {
	//Validation for User Inputs:
	isValidNmae := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= reamainingTicket

	return isValidNmae, isValidEmail, isValidTicketNumber
}
