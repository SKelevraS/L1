package main

import "fmt"

type Human struct {
	Name   string
	Gender string
	hobby string
}

type MoreInfo struct {
	Height float32
	Weight float32
	Human
foMoreInfo}

func main() {
	// Вложили в переменную MyHuman информацию по человеку с именем Dan
	MyHuman := MoreInfo{
		Height: 180,
		Weight: 62,
		Human: Human{
			Name:   "Dan",
			Age:    19,
			Hobby: "Programming"}}

	// Выводим данные переменной на экран обратившись к переменной
	fmt.Println(MyHuman)

	// Выводим данные переменной на экран обратившись к переменной с определенным полем
	fmt.Println(MyHuman.Name, MyHuman.Age, MyHuman.Height)
}