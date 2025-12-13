package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    // The strings.Repeat function is used to create the border line.
    borderLine := strings.Repeat("*", numStarsPerLine)
    
    return borderLine + "\n" + welcomeMsg + "\n" + borderLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    // 1. Remove all stars using strings.ReplaceAll.
    noStars := strings.ReplaceAll(oldMsg, "*", "")
    
    // 2. Remove leading and trailing whitespace using strings.TrimSpace.
    return strings.TrimSpace(noStars)
}