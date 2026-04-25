package main

import "fmt"

const course = 1 / 5828215.00 //на 26.04.2026
const negativeMoney = "Ошибка: Сумма платежа должна быть положительной"
const notEnoughMoney = "Ошибка: У Вас не достаточно средств."

type PaymentProcessor interface {
	GetParams() map[string]string
	Process(amount float64) string
}

type CreditCard struct {
	PaySystem  string
	CardNumber string
	Currency   string
	Balance    float64
}

func (account *CreditCard) GetParams() map[string]string {
	return map[string]string{
		"system":  account.PaySystem,
		"id":      account.CardNumber,
		"balance": fmt.Sprintf("%.2f RUB", account.Balance)}
}

func (account *CreditCard) Process(amount float64) string {
	if amount <= 0 {
		return fmt.Sprintf(negativeMoney)
	}
	if account.Balance < amount {
		return fmt.Sprintf(notEnoughMoney)
	}
	account.Balance -= amount
	return fmt.Sprintf("Оплата картой #%s на сумму %.2f руб. успешно обработана. Баланс карты: %.2f",
		account.CardNumber, amount, account.Balance)
}

type CryptoWallet struct {
	PaySystem string
	WalletId  string
	Currency  string
	Balance   float64
}

func (account *CryptoWallet) GetParams() map[string]string {
	return map[string]string{
		"system":  account.PaySystem,
		"id":      account.WalletId,
		"balance": fmt.Sprintf("%e BTC", account.Balance),
	}
}

func (account *CryptoWallet) Process(amount float64) string {
	if amount <= 0 {
		return fmt.Sprintf(negativeMoney)
	}
	cryptAmount := amount * course
	if account.Balance < cryptAmount {
		return fmt.Sprintf(notEnoughMoney)
	}
	account.Balance -= cryptAmount
	return fmt.Sprintf("Оплата кошельком #%s на сумму %.2f RUB. (%.e %s) успешно обработана. Баланс: %e",
		account.WalletId, amount, cryptAmount, account.Currency, account.Balance)
}

var accounts = []PaymentProcessor{
	&CreditCard{
		PaySystem:  "Банковская карта",
		CardNumber: "1111 2222 3333 4444",
		Balance:    5000,
		Currency:   "RUB",
	},
	&CryptoWallet{
		PaySystem: "Виртуальный кошелек",
		WalletId:  "987654321987654321",
		Balance:   20000.00 * course,
		Currency:  "BTC",
	},
	&CreditCard{
		PaySystem:  "Банковская карта",
		CardNumber: "2222 3333 4444 5555",
		Balance:    10000,
		Currency:   "RUB",
	},
	&CryptoWallet{
		PaySystem: "Виртуальный кошелек",
		WalletId:  "123456789123456789",
		Balance:   5000.00 * course,
		Currency:  "BTC",
	},
}

func main() {
	runApp()
}

func runApp() {
	for _, processor := range accounts {
		fmt.Printf("%s #%s:\n%s\n\n", processor.GetParams()["system"],
			processor.GetParams()["id"], processor.GetParams()["balance"])
		amounts := []float64{500.00, -500.00, 1200.00, 75.50, 5000.00}
		for _, amount := range amounts {
			fmt.Printf("Платёж %.2f RUB: %s\n", amount, processor.Process(amount))
		}
		fmt.Println()
	}
}
