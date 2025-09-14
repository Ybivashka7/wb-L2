package main

import (
	"fmt"
	"os"
)

// Данная функция вовращает интерфейс с типом *os.PathError и значение nil , поэтому при выводе мы видим nil
// а при проверке на nil false  потому что интерфейсы в go  хранят тип и значение , а так как у err есть тип , но значение nil
// то при проверке на nil  мы видим false потому что интерфейс не пустой , в нем есть тип данных
func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
