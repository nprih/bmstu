package main

import "fmt"

//break, continue

func main() {
	age := 0
	for {
		fmt.Printf("Укажите возраст: ")
		fmt.Scanln(&age)
		res := "Стоимость проезда: "
		if age <= 0 || age > 116 {
			res = "Вы врете"
		} else if age >= 65 || age <= 7 {
			res += "бесплатно"
		} else {
			res += "30 рублей"
		}
		fmt.Println(res)
	}
}
