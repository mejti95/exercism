package booking

import ("time"
        s "strconv")
 

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/02/2006 15:04:05"
    scheaduleLayout, _ := time.Parse(layout, date)
    scheaduleLayout.Format("2006-01-02T15:04:05.0000.UTC")
	return scheaduleLayout
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January  2, 2006 15:04:05"
    dateToCheck, _ := time.Parse(layout, date)
    timeNow := time.Now()
    return timeNow.After(dateToCheck)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    dateToCheck, _ := time.Parse(layout, date)
    if dateToCheck.Hour() >= 12 &&  dateToCheck.Hour() < 18 {
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    dateToCheck, _ := time.Parse(layout, date)
    
    return "You have an appointment on " + dateToCheck.Weekday().String() + ", " + dateToCheck.Month().String() + " " + s.Itoa(dateToCheck.Day()) + ", " + s.Itoa(dateToCheck.Year()) + ", at " + s.Itoa(dateToCheck.Hour()) + ":" + s.Itoa(dateToCheck.Minute()) + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now()
    t := time.Date(currentYear.Year(),time.September,15,00,0,0,0,time.UTC)
    return t
}
