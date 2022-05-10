package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, -9.0, 15.5, 24.5, -21.0, 32.5, -5.0, 5.0}
	endMap := make(map[int][]float64)

	for _, temp := range arr {
		if temp < 0 {
			mapKey := (int(temp) - 10) / 10 * 10 //Округляем элементы среза и получаем 'группу', в которую будут поступать значения, если они меньше 0
			endMap[mapKey] = append(endMap[mapKey], temp)

		} else {
			key := (int(temp) / 10) * 10 //Повторяем тоже самое с числами больше нуля
			endMap[key] = append(endMap[key], temp)
		}
	}

	fmt.Println(endMap)

}
