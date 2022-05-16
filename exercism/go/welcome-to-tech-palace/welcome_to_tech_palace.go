package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    msg := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
    return msg
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    msg := 
    strings.Repeat("*", numStarsPerLine) + "\n" + 
    welcomeMsg + "\n" + 
    strings.Repeat("*", numStarsPerLine)
    return msg
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    rmBacktick := strings.ReplaceAll(oldMsg, "`", "")
    rmAsterisks := strings.ReplaceAll(rmBacktick, "*", "")
    cleanMsg := strings.TrimSpace(rmAsterisks)
    return cleanMsg
	panic("Please implement the CleanupMessage() function")
}
