package helper

import "strings"

// Variables and Functions defined outside any function, can
// be accessed in all other files within the same package.
// To export a function, it must be capitalized e.g., public, private, etc.

func ValidateUserInput(firstName string, lastName string, userEmail string, userOrders uint, remainingRendang uint) (bool, bool, bool) {
	// Validation checks (|| is OR operator, != is NOT operator)
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidOrders := userOrders > 0 && userOrders <= remainingRendang
	return isValidName, isValidEmail, isValidOrders
}
