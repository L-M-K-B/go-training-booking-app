package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")
}
