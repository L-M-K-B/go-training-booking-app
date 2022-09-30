package main

import (
	"fmt"
	"strings"
)

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50

	// bookings as an array
	// var bookingsArray [50]string
	// or as a slice
	var bookingsSlice []string

	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")

	for {
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

		remainingTickets = remainingTickets - userTickets
		//bookingsArray[0] = firstName + " " + lastName

		//fmt.Printf("The whole array: %v\n", bookingsArray)
		//fmt.Printf("The first value: %v\n", bookingsArray[0])
		//fmt.Printf("Array type: %T\n", bookingsArray)
		//fmt.Printf("Array length: %v\n", len(bookingsArray))

		bookingsSlice = append(bookingsSlice, firstName+" "+lastName)
		fmt.Printf("The whole slice: %v\n", bookingsSlice)
		//fmt.Printf("The first value: %v\n", bookingsSlice[0])

		fmt.Printf("Thanks %v %v for booking %v tickets.\nYou will receive a confirmation at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, confName)

		var firstNames []string
		for _, singleBooking := range bookingsSlice {
			var names = strings.Fields(singleBooking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		fmt.Println("\n")
	}
}
