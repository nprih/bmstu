package main

import "fmt"

type celsius float64
type kelvin float64
type farenheit float64

func (t farenheit) celsius() celsius {
	//Без логики
	return 100
}

func (t kelvin) celsius() celsius {
	t -= 100
	return celsius(t - 273.15)
}

func main() {
	var temperatureInCelsius celsius = 10.7
	var temperatureInKelvin kelvin = 400
	var temperatureInFahrenheit farenheit = 10
	fmt.Println(temperatureInCelsius + temperatureInKelvin.celsius())
	fmt.Println(temperatureInCelsius + temperatureInFahrenheit.celsius())
}
