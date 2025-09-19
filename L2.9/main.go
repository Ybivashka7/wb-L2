package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	fmt.Println(unpakingString(str))
}
func unpakingString(str string) (string, error) {
	//если строка пустая, то возвращаем nil и ошибку
	if str == "" {
		return "", fmt.Errorf("пустая строка")
	}
	// Builder для запоминания символов
	var bul strings.Builder
	var prev rune
	escape := false
	// в цикле бежим по строке и проверяем каждый символ
	for _, v := range str {
		switch {
		// экранированный символ пишем как есть
		case escape:
			bul.WriteRune(v)
			prev = v
			escape = false
			// отмечаем  экранирование
		case v == '\\':
			escape = true
			// проверяем на цифру и добавляем предыдущий символ такое же количество раз как и цифра
		case unicode.IsDigit(v):
			if prev == 0 {
				return "", fmt.Errorf("некоректная строка+")
			}
			count := int(v - '0')
			if count == 0 {
				prev = 0
				continue
			}
			for i := 1; i < count; i++ {
				bul.WriteRune(prev)
			}
			// просто записываем символ
		default:
			bul.WriteRune(v)
			prev = v

		}
	}
	return bul.String(), nil
}
