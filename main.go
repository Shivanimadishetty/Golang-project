package main

import "fmt"

func main() {
	var exhibitionName = "Art Exhibition"
	const totalTickets = "50"
	var remainingTickets = "50"

	fmt.Printf("Welcome to my %v Hope you'll have a great time \n", exhibitionName)
	fmt.Printf("We have a total of %v tickets available and the number of tickets remaining are %v \n", totalTickets, remainingTickets)
	fmt.Println("Book your tickets here")

}
