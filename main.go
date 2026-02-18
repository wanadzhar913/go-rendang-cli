package main

import (
	"fmt"
	"strings"
)

func main() {
	// var eventName = "Rendang Factory"
	eventName := "Rendang Factory" // short declaration operator (for non-const variables)

	const rendangStock uint = 100
	var remainingRendang uint = 100

	// Arrays in Go have a fixed size
	// var bookings [50]string // it's a fixed size array of 50 elements

	// Slices in Go are dynamic arrays
	var bookings []string // it's a dynamic array of strings

	// Print "Hello, World!" (with a newline) to the console
	fmt.Printf("Welcome to our %v!\n", eventName)
	fmt.Println("Get your Rendang stock here!")
	fmt.Printf("We have %.2f packets of rendang available.\n", float64(remainingRendang))

	// Infinite loop (with conditions)
	for remainingRendang > 0 && len(bookings) < 100 {
		var firstName string
		var lastName string
		var userEmail string
		var userOrders uint

		// Ask for user input
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName) // & is a 'Pointer' and it's a variable that points to the memory addres of another variable

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter your orders:")
		fmt.Scan(&userOrders)

		// Validation checks (|| is OR operator, != is NOT operator)
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")

		if isValidName && isValidEmail {

			if userOrders > remainingRendang {
				fmt.Printf("Sorry, we only have %.2f packets of rendang left.\n", float64(remainingRendang))
				continue // continue to the next iteration of the loop
			} else if len(bookings) == 100 {
				fmt.Println("Sorry, we've reached the maximum number of bookings.")
				break // break out of the loop
			}

			remainingRendang -= userOrders
			// bookings[0] = firstName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Hello %v, you have ordered %v packets of rendang. ", firstName, userOrders)
			fmt.Printf("We'll send a confirmation email to %v.\n\n", userEmail)

			fmt.Printf("We've now only %.2f packets of rendang left. ", float64(remainingRendang))
			fmt.Printf("We have %v booking(s)\n\n", len(bookings))

			firstNames := []string{}

			// for-each loop
			// think of _ (blank identifier) as Python's enumerate() function
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0]) // append() returns a new slice with the new element added
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Println("Thank you for your order!")

			noRendangRemaining := remainingRendang == 0 // boolean
			if noRendangRemaining {
				fmt.Println("Our rendang is sold out! See you next Ramadan & Selamat Hari Raya!")
				break // break out of the loop
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name. It's too short!")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address.")
			}
			continue // continue to the next iteration of the loop
		}
	}

	city := "KL"
	// switch statements allow for variables to be tested for equality against a list of values
	switch city {
	case "Malaysia", "KL":
		fmt.Println("You're in Malaysia.")
	case "Singapore":
		fmt.Println("You're in Singapore.")
	case "Indonesia", "Jakarta":
		fmt.Println("You're in Indonesia.")
	}
}
