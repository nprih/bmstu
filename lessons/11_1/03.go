package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
С помощью цикла for запустить 5 горутин
			И любыми известными методами, правдами и неправдами
			Заставить их запуститься точно по очереди
	(Проверяется выводом индекса горутины на экран)
*/

func main() {
	str := ""
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			str += strconv.Itoa(index) + "\n"
			wg.Done()
		}(i)
		wg.Wait()
	}
	fmt.Println(str)
}
