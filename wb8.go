
package main

import (
	"fmt"
	"strconv"
)

func main() {
	dataMap := make(map[int]int64)

	var helpScan int64

	// Заполняем карту введенными данными с клавиатуры, присваивая i в качестве ключа
	for i := 0; i < 3; i++ {
		_, err := fmt.Scan(&helpScan)
		checkError(err)
		if i == 1 && (helpScan < 1 || helpScan > 64){
			fmt.Println("Нет")
		}
		if i == 2 && (helpScan != 0 && helpScan != 1){
			fmt.Println("Тоже нет")
		}
		dataMap[i] = helpScan
	}

	fmt.Println(strconv.FormatInt(dataMap[0], 2))

	var endNum int64

	if dataMap[1] == 64{ //Обработка битового знака
		fmt.Println(strconv.FormatInt(-dataMap[0], 2))
		return
		
	}

	if dataMap[2] == 1 {
		endNum = dataMap[0] | dataMap[2]<<(dataMap[1]-1)
	} else if dataMap[2] == 0 {
		endNum = dataMap[0] & (9223372036854775807 - 1<<(dataMap[1]-1)) //?
	}

	fmt.Println(strconv.FormatInt(endNum, 2))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}