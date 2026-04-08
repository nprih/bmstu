package main

import (
	"fmt"
	"time"
)

func main() {
	// 01   02   03   04   05   06   Mon
	// Jan 02.  15.   04.  05.  2006 Monday
	// 2 ярваря 2006 15 часов: 04 минуты: 05 секунд

	nowTime := time.Now()
	//fmt.Println(nowTime)
	//fmt.Println(nowTime.Format("2006 Jan 02 03:04:05 Monday"))
	//fmt.Println(nowTime.Format("02.01.06"))
	//fmt.Println(nowTime.Format(time.RFC3339))
	//fmt.Println(nowTime.Format(time.DateTime))
	//fmt.Println(nowTime.Format("06.02.01 15:04"))
	//fmt.Println(nowTime.Format("02 January 06"))
	fmt.Println(nowTime.Format("15:04 Mon 2006/01/02"))
}
