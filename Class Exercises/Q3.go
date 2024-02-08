/*
You're developing a weather monitoring system in GoLang for a research institute.
The system needs to store data about temperature, humidity, and wind speed. Define variables to hold these measurements for a specific location at a given point in time.
Ensure you choose suitable data types for representing numerical measurements accurately.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// I have created a struct for weather measurements
	type Weather struct {
		Time        time.Time
		Temperature float64
		Humidity    float64
		WindSpeed   float64
	}

	// I have created an empty slice to store weather data for multiple locations
	var locations []Weather

	// I am taking user inputs for the number of locations
	var numLocations int
	fmt.Print("Enter the number of locations: ")
	fmt.Scanln(&numLocations)

	// I'm iterating over each location to input weather data
	for i := 0; i < numLocations; i++ {
		var (
			temperature float64
			humidity    float64
			windSpeed   float64
		)

		fmt.Printf("\nEnter weather details for Location %d:\n", i+1)

		fmt.Print("Temperature (in Celsius): ")
		fmt.Scanln(&temperature)

		fmt.Print("Humidity (in percentage): ")
		fmt.Scanln(&humidity)

		fmt.Print("Wind Speed (in meters per second): ")
		fmt.Scanln(&windSpeed)

		// Here current time is recorded
		currentTime := time.Now()

		// Here I have created a new Weather object and append it to the slice
		locationWeather := Weather{
			Time:        currentTime,
			Temperature: temperature,
			Humidity:    humidity,
			WindSpeed:   windSpeed,
		}
		locations = append(locations, locationWeather)
	}

	// Here I am displaying the weather data for each location
	fmt.Println("\nWeather Data for Each Location:")
	for i, location := range locations {
		fmt.Printf("Location %d (Recorded at %s):\n", i+1, location.Time.Format("2006-01-02 15:04:05"))
		fmt.Printf("Temperature: %.2f°C\n", location.Temperature)
		fmt.Printf("Humidity: %.2f%%\n", location.Humidity)
		fmt.Printf("Wind Speed: %.2f m/s\n", location.WindSpeed)
		fmt.Println()
	}
}
