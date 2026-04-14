package main

import "fmt"

type celsius float64
type kelvin float64

//func kelvinToCelsius(t kelvin) celsius {
//	return celsius(t - 273.15)
//}

func (t kelvin) сelsius() celsius {
	return celsius(t - 273.15)
}

func main() {

	var temperatureInCelsius celsius = 10.7
	var temperatureInKelvin kelvin = 400

	fmt.Println(temperatureInCelsius + temperatureInKelvin.сelsius())
}
