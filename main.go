package main

import "fmt"

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")

	// Book tickets
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thanks %v %v for booking %v tickets.\n You will receive a confirmation at %v.\n", firstName, lastName, userTickets, email)
}
