// Package weather provides tools to forecast the weather.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string
// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast returns a string represents the current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
