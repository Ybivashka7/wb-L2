package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция  принимает числа и возвращает канал
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			// случайная задержка
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c) //  закрываем канал
	}()
	return c
}

// merge — объединяет два канала в один (fan-in)
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a: // читаем из первого канала
				if ok {
					c <- v
				} else {
					a = nil // если канал закрыт — исключаем из select
				}
			case v, ok := <-b: // читаем из второго канала
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			// когда оба канала закрыты — закрываем все
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())

	a := asChan(1, 3, 5, 7) // первый канал
	b := asChan(2, 4, 6, 8) // второй канал

	c := merge(a, b) // слияние каналов

	// читаем из одного канала
	for v := range c {
		fmt.Print(v)
	}
	fmt.Println()
}
