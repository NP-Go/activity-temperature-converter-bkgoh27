package main

import "fmt"

func convertKelvin(temperatureInput float64) (float64, float64) {
	var resultFahrenheit = 9.0 / 5.0 * (temperatureInput - 459.67)
	var resultCelsius = 5.0 / 9.0 * (resultFahrenheit - 32.0)
	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	var resultFahrenheit = 9.0 / 5.0 * (temperatureInput + 32.0)
	var resultKelvin = 5.0 / 9.0 * (resultFahrenheit + 459.67)
	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	var resultKelvin = 5.0 / 9.0 * (temperatureInput + 459.67)
	var resultCelsius = 5.0 / 9.0 * (temperatureInput - 32.0)
	//Do not remove this line
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput, resultFahrenheit, resultCelsius, resultKelvin float64
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius and 3 for Fahrenheit.): ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		resultFahrenheit, resultCelsius = convertKelvin(temperatureInput)
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Celsius: ", resultCelsius)
	} else if temperatureChoice == 2 {
		resultFahrenheit, resultKelvin = convertCelsius(temperatureInput)
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Kelvin: ", resultKelvin)
	} else if temperatureChoice == 3 {
		resultKelvin, resultCelsius = convertFahrenheit(temperatureInput)
		//DO not remove this
		fmt.Println("Kelvin: ", resultKelvin, " Celsius: ", resultCelsius)
	} else {
		fmt.Println("No Conversion")
	}

}
