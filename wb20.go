package main

import "fmt"

func main() {
	message := "snow dog sun"

	messageSlice := make(map[int]string) //Лучше слайс

	num := 0
	var word string

	for _, i := range message {

		if string(i) != " " {
			word += string(i)
		} else {
			messageSlice[num] = word
			fmt.Println(messageSlice[num])
			word = ""
			num++
		}
	}
	messageSlice[num] = word // Записываем последнее слово в карту

	var endMessage string

	// Перебираем все значения в карте и прибавляем их к новой переменной, так же не забываем прибавить пробел
	for i := 0; i < len(messageSlice); i++ {
		endMessage = messageSlice[i] + " " + endMessage
	}

	fmt.Println(endMessage)
}