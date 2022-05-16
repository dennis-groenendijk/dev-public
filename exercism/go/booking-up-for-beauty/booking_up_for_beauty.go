package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    t, _ := time.Parse("1/2/2006 15:04:05", date)
    return t
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    t, _ := time.Parse("January 2, 2006 15:04:05", date)
    return t.Before(time.Now())
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    return t.Hour() >= 12 && t.Hour() <= 18
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    t := Schedule(date)
    return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    now := time.Now()
    return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	panic("Please implement the AnniversaryDate function")
}
