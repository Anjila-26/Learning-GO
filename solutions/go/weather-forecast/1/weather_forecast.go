// Package weather implements the routine of fetching and displaying the weather information.
package weather

var (
    // CurrentCondition represents the condition that is happening in the present.
	CurrentCondition string

    // CurrentLocation represents the location for which we are looking the condition and weather.
	CurrentLocation  string
)

// Forecast function returns us the current condition for the current location we have.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
