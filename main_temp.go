package main

import "fmt"

func main() {

	var exhibitionName = "Art Exhibition"
	const totalTickets = "50"
	var remainingTickets uint = 50
	var Bookings [50]string

	fmt.Printf("Welcome to my %v Hope you'll have a great time \n", exhibitionName)
	fmt.Printf("We have a total of %v tickets available and the number of tickets remaining are %v \n", totalTickets, remainingTickets)
	fmt.Println("Book your tickets here")

	var FirstName string
	var LastName string
	var Email string
	var userTickets uint

	fmt.Printf("Please enter your First name: ")
	fmt.Scanln(&FirstName)

	fmt.Printf("Please enter your Last Name: ")
	fmt.Scanln(&LastName)

	fmt.Printf("Please enter your Email address: ")
	fmt.Scanln(&Email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	remainingTickets = remainingTickets - userTickets
	Bookings[0] = FirstName + " " + LastName

	fmt.Printf("\nThank you %v %v for booking %v tickets! \n", FirstName, LastName, userTickets)
	fmt.Printf("You will receive a confirmation on your email address %v \n", Email)
	fmt.Printf("%v tickets remaing for %v \n", remainingTickets, exhibitionName)
	fmt.Printf("The whole array %v: \n", Bookings)
	fmt.Printf("The First value %v: \n", Bookings[0])
	fmt.Printf("The array type %T: \n", Bookings)
	fmt.Printf("The array length %v: \n", len(Bookings))
}
