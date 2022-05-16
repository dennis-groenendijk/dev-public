// Package weather returns the forecast based on a particular city and weather condition.
package weather

// CurrentCondition states the current weather condition.
var CurrentCondition string

// CurrentLocation states the current location.
var CurrentLocation string

// Forecast returns a string which displays the weather condition of the location that is put in.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
