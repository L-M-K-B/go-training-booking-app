package main

import (
	"fmt"
	"sync"
	"time"
)

const confTickets = 50

var confName = "Go Conference"
var remainingTickets uint = 50
var bookingsSlice = make([]UserData, 0) //=> 0 is the initial number of elements in the map, will be updated automatically by adding new elements

type UserData struct {
	firstName    string
	lastName     string
	email        string
	ticketNumber uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		// Book tickets
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			printFirstNames(bookingsSlice)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name is to short. Try again!")
			}
			if !isValidEmail {
				fmt.Println("The email address doesn't contain an @. Try again!")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets is invalid. Try again!")
			}
			fmt.Println("")
			continue
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")
}

func printFirstNames(bookingsSlice []UserData) {
	var firstNames []string
	for _, singleBooking := range bookingsSlice {
		firstNames = append(firstNames, singleBooking.firstName)
	}
	//fmt.Printf("The first names of bookings are: %v\n", firstNames)
	//fmt.Println("\n")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ignore error handling for Scan
	fmt.Println("Enter your first name:")
	_, _ = fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	_, _ = fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	_, _ = fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	_, _ = fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		ticketNumber: userTickets,
	}

	bookingsSlice = append(bookingsSlice, userData)

	fmt.Printf("Thanks %v %v for booking %v tickets.\nYou will receive a confirmation at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, confName)
	fmt.Println("")

	return remainingTickets, bookingsSlice
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: \n%v \nto email address%v\n", ticket, email)
	fmt.Println("####################")

	wg.Done()
}
