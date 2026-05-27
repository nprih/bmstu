package input

import (
	"2_7/functions"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pare struct {
	Login    string
	Password string
}

func InputAuth() (error, Pare) {
	fmt.Print("Введите логин(email): ")
	reader := bufio.NewReader(os.Stdin)
	login, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return err, Pare{}
	}
	login = functions.ToLowerCase(strings.TrimSpace(login))

	fmt.Print("Введите пароль: ")
	pass, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return err, Pare{}
	}
	pass = strings.TrimSpace(pass)

	return nil, Pare{login, pass}
}
