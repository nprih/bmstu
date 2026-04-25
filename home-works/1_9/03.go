package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (account BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Ошибка: сумма должна быть положительной\n")
		return
	}
	account.Balance += amount
	fmt.Printf("Счет пополнен на %.2f. Текущий баланс: %.2f\n\n", amount, account.Balance)
}

func (account BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Ошибка: сумма должна быть положительной\n")
		return
	}
	if amount > account.Balance {
		fmt.Printf("Ошибка: недостаточно средств. Доступно: %.2f, запрошено: %.2f\n\n", account.Balance, amount)
		return
	}
	account.Balance -= amount
	fmt.Printf("Снято %.2f. Текущий баланс: %.2f\n\n", amount, account.Balance)
}

func main() {
	account := BankAccount{
		Owner:   "Алёша",
		Balance: 1000,
	}

	fmt.Printf("Создан счет, владелец %s, баланс %.2f\n\n", account.Owner, account.Balance)

	runApp(account)

}

func runApp(account BankAccount) {
	fmt.Println("Попытка пополнения на 500:")
	account.Deposit(500.00)
	fmt.Println("Попытка пополнения на -500:")
	account.Deposit(-500.00)
	fmt.Println("Попытка снятия  300:")
	account.Withdraw(300.00)
	fmt.Println("Попытка снятия 1500:")
	account.Withdraw(1500.00)
	fmt.Println("Попытка снятия -100:")
	account.Withdraw(-100.00)
}
