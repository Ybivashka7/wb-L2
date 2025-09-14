package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	// используем NPT библиотеку
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	// Обрабатываем ошибку
	if err != nil {
		fmt.Println("Ошибка получения времени ", err)
	}
	// выводим полученное время
	fmt.Println(time)
}
