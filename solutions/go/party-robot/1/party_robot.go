package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    // string(age) converts the number to a Unicode character, not the digits.
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
        // %s is for strings (name, direction, neighbor).
    // %02d is for the integer table number (ensures a leading zero, e.g., 5 becomes 05).
    // %.2f is for the float distance (ensures two decimal places).
    
	return fmt.Sprintf("%s\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", 
		Welcome(name), 
		table, 
		direction, 
		distance, 
		neighbor,
	)
}