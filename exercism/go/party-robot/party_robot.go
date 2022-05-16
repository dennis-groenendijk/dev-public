package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    w := fmt.Sprintf("Welcome to my party, %s!", name)
    return w
	panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    hb := fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
    return hb
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
// The print statement should be one line, but is broken up for better readability.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    at := fmt.Sprintf("Welcome to my party, %s!\n" + 
                      "You have been assigned to table %03d. Your table is %s, " +
                      "exactly %.1f meters from here.\nYou will be sitting next to %s.", name, 							  table, direction, distance, neighbor)
    return at
	panic("Please implement the AssignTable function")
}
