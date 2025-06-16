package main

import (
	"bookingapp/valid"
	"fmt"
	"sync"
	"time"
)

var ConfName string = "Go Conference"

const Tickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	FirstName     string
	LastName      string
	Email         string
	BookedTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	var RemainingTickets uint = 50

	// This function greets the user and displays the conference details
	GreetUser()

	for {
		// This function prompts the user to enter their details and returns the values
		fName, lName, Email, bookedTickets := EnterUserDetails()

		// Validate user input

		isValidName, isValidEmail, isValidTicketNumber := valid.ValidUserInput(fName, lName, Email, bookedTickets, RemainingTickets)

		if !isValidName {

			fmt.Println("The name you entered is TOO SHORT. Please try again.")
			continue
		}
		if !isValidEmail {
			fmt.Println("The email you entered is INVALID. Please try again.")
			continue
		}
		if !isValidTicketNumber {
			fmt.Println("You have entered an INVALID number of tickets. Please try again.", RemainingTickets)
			continue
		}

		RemainingTickets = BookTickets(RemainingTickets, bookedTickets, fName, lName, Email)

		wg.Add(1)
		go SendTicket(fName, lName, Email, bookedTickets)

		if RemainingTickets == 0 {
			fmt.Printf("All tickets have been SOLD OUT.")
			break
		}
		//end program
	}
	wg.Wait()
	fmt.Println("\n🎉 Thank you for booking with Go Conference!")
	fmt.Println("We're super excited to have you on board. 😊")
	fmt.Println("Check your inbox for your tickets — see you at the event! 🥳")
}

func GreetUser() {
	fmt.Println("Welcome to", ConfName, ". We are thrilled to have you join us for this exciting event.")
	fmt.Println("We have a total of", Tickets, "tickets available.")
}

func EnterUserDetails() (string, string, string, uint) {

	var fName, lName, Email string
	var bookedTickets uint

	fmt.Println("Enter FIRST NAME: ")
	fmt.Scan(&fName)

	fmt.Println("Enter LAST NAME: ")
	fmt.Scan(&lName)

	fmt.Println("Enter EMAIL ADDRESS: ")
	fmt.Scan(&Email)

	fmt.Println("Number of Tickets: ")
	fmt.Scan(&bookedTickets)
	return fName, lName, Email, bookedTickets
}
func BookTickets(RemainingTickets uint, bookedTickets uint, fName string, lName string, Email string) uint {
	// This function books the tickets and updates the remaining tickets

	userData := UserData{

		FirstName:     fName,
		LastName:      lName,
		Email:         Email,
		BookedTickets: bookedTickets,
	}
	fmt.Printf("\n✅ Thank you %v %v for booking %v ticket(s).\n", fName, lName, bookedTickets)
	fmt.Printf("A confirmation has been sent to %v.\n", Email)
	fmt.Printf("🎟️ Tickets remaining: %v\n\n", RemainingTickets-bookedTickets)

	// Append the user data to the bookings slice
	bookings = append(bookings, userData)
	fmt.Println(bookings)
	return RemainingTickets - bookedTickets
}
func SendTicket(fName string, lName string, Email string, bookedTickets uint) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", bookedTickets, fName, lName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, Email)
	fmt.Println("#################")
	wg.Done()
}
