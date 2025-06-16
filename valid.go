package valid

import (
	"strings"
)

func ValidUserInput(fName string, lName string, Email string, bookedTickets uint, RemainingTickets uint) (bool, bool, bool) {

	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(Email, "@") && strings.Contains(Email, ".")
	isValidTicketNumber := bookedTickets <= RemainingTickets && bookedTickets > 0
	return isValidName, isValidEmail, isValidTicketNumber

}
