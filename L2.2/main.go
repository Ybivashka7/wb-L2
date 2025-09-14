package main

import "fmt"

// В данной функции мы делаем именное возвращение (х), поэтому перед возвратом функция defer повлияет на переменную х
func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

// В этой функции у нас нет именного возвращаемого значения поэтому функция defer не повляем на переменную x
// так как именной переменной нет то return запомнит x=1 и defer никак не повлияет на значение
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
