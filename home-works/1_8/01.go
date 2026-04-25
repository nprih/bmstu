package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dividend, divisor, err := strToFloat()
	if err != nil {
		return
	}

	res, dividingByZero := divide(dividend, divisor)
	if dividingByZero != nil {
		fmt.Printf("Ошибка: %s!\n", dividingByZero)
		return
	}
	fmt.Println(res)
}

func inputFloat() (dividend string, divisor string, err error) {
	fmt.Print("Введите делимое: ")
	reader := bufio.NewReader(os.Stdin)
	text1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Введите делитель: ")
	text2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	return strings.TrimSpace(text1), strings.TrimSpace(text2), err
}

func strToFloat() (dividend float64, divisor float64, err error) {
	str1, str2, err := inputFloat()
	if err != nil {
		return
	}
	dividend, err = strconv.ParseFloat(strings.ReplaceAll(str1, ",", "."), 64)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	divisor, err = strconv.ParseFloat(strings.ReplaceAll(str2, ",", "."), 64)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	return dividend, divisor, err
}

func divide(dividend float64, divisor float64) (result float64, err error) {
	if divisor == 0 {
		return 0, errors.New("на ноль делить нельзя, так сказал калькулятор")
	}
	return dividend / divisor, nil
}
