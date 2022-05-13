package main

import (
	"fmt"
	"time"
)

func main() {
	// func Now() Time
	// возвращает текущую дату и время
	now := time.Now()

	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// возвращает дату и время в соответствии с заданными параметрами: годом, месяцем, днем, временем и пр.
	currentTime := time.Date(
		2020,     // год
		time.May, // месяц
		15,       // день
		10,       // часы
		13,       // минуты
		12,       // секунды
		45,       // наносекунды
		time.UTC, // временная зона
	)

	// func Unix(sec int64, nsec int64) Time
	// возвращает дату и время в соответствии с заданными параметрами: секундами и наносекундами, прошедшими с начала эпохи Unix — 01.01.1970 г.
	// https://ru.wikipedia.org/wiki/Unix-%D0%B2%D1%80%D0%B5%D0%BC%D1%8F
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)

	fmt.Println(now.Format("02-01-2006 15:04:05"))         // 15-05-2020 09:58:16
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00
}
