package main

import (
	"fmt"
	"go-rendang-cli/helper"
	"sync"
	"time"
)

// eventName := "Rendang Factory" // short declaration operator (for non-const variables), usable only in function scope
var eventName = "Rendang Factory"

const rendangStock uint = 100 // uint is an unsigned integer (positive integers)

var remainingRendang uint = 100

// Arrays in Go have a fixed size
// var bookings [50]string // it's a fixed size array of 50 elements (no mixed types!)
// Slices in Go are dynamic arrays (and we don't have to pay attention to the index)
// var bookings []string // it's a dynamic array of strings (no mixed types!)
// var bookings = []string{}
// Maps are key-value pairs (they support only 1 data type)
var bookings = make([]UserData, 0) // create empty slice of UserData. The initial size is 0.

// Structs (structures) are collections of fields (they support multiple data types)
type UserData struct {
	firstName  string
	lastName   string
	userEmail  string
	userOrders uint
}

var wg = sync.WaitGroup{} // WaitGroup is used to wait for all goroutines to finish.
var mu sync.Mutex         // protects shared state (e.g., remainingRendang). A Mutex is a mutual exclusion lock that allows only one goroutine to access the shared state at a time.

func main() {

	greetUser()

	for len(bookings) < 100 {

		firstName, lastName, userEmail, userOrders := getUserInput()

		mu.Lock()
		isValidName, isValidEmail, isValidOrders := helper.ValidateUserInput(firstName, lastName, userEmail, userOrders, remainingRendang)
		mu.Unlock()

		if isValidName && isValidEmail && isValidOrders {

			mu.Lock()
			bookRendang(firstName, lastName, userEmail, userOrders)
			mu.Unlock()

			wg.Add(1)
			go sendConfirmationEmail(userOrders, firstName, lastName, userEmail) // go keyword is used to create a new goroutine (a lightweight thread managed by the Go runtime). It's asynchronous.

			firstNames := printFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Println("Thank you for your order!")

			noRendangRemaining := remainingRendang == 0 // boolean
			if noRendangRemaining {
				fmt.Println("Our rendang is sold out! See you next Ramadan & Selamat Hari Raya!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name. It's too short!")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address.")
			}
			if !isValidOrders {
				fmt.Println("Please enter a valid number of orders.")
			}
			if len(bookings) == 100 {
				fmt.Println("Sorry, we've reached the maximum number of bookings.")
				break
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

	wg.Wait()
	fmt.Printf("All confirmation emails have been sent. List of all bookings:\n%v\n", bookings)
}

func greetUser() {
	fmt.Printf("Welcome to our %v!\n", eventName)
	fmt.Println("Get your Rendang stock here!")
	fmt.Printf("We have %.2f packets of rendang available.\n", float64(remainingRendang))
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var orders uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) // & is a 'Pointer' and it's a variable that points to the memory addres of another variable

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your orders:")
	fmt.Scan(&orders)

	return firstName, lastName, email, orders
}

func printFirstNames() []string {
	// for-each loop
	// think of _ (blank identifier) as Python's enumerate() function
	firstNames := []string{} // we use {} when we want a non-nil slice

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName) // append() returns a new slice with the new element added
	}
	return firstNames
}

func bookRendang(firstName string, lastName string, email string, orders uint) {
	remainingRendang -= orders

	// create user map
	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		userEmail:  email,
		userOrders: orders,
	}
	bookings = append(bookings, userData)
}

func sendConfirmationEmail(orders uint, firstName string, lastName string, email string) {
	defer wg.Done()

	time.Sleep(10 * time.Second) // simulate email sending time (crucially, it blocks the current "thread" (goroutine) execution for the defined duration). Only when it's done, the next line of code will be executed.

	var orderSummary = fmt.Sprintf("%v %v has ordered %v packets of rendang.", firstName, lastName, orders)
	fmt.Println("###########")
	fmt.Printf("Sending confirmation email: %v\nTo: %v...\n", orderSummary, email)
	fmt.Println("###########")

	// send email
	// TODO: implement email sending functionality
}
