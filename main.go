package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)
	// Check ticket counts
	if isValidTickets && isValidEmail && isValidName {
		bookTickets(userTickets, firstName, lastName, email)

		// Concurrency for send mail section
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstNames()
		fmt.Printf("First names of attendees: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets sold out!")
		}

	} else {

		if !isValidName {
			fmt.Println("First name or last names is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address not valid")
		}
		if !isValidTickets {
			fmt.Println("Tickets required is invalid")
		}
	}
	wg.Wait()
}

// Functions

func greetUser() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and, %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	// Get attendees first names
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// Get user input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// get user input
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Please enter your desired ticket Qty:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	// Book user ticket

	remainingTickets = remainingTickets - userTickets

	// Mapping user info
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n%v \nto: %v\n", ticket, email)
	fmt.Println("############")
	wg.Done()
}
