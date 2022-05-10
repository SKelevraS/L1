package main

import "fmt"

type Human struct {
	Name   string
	Hobby  string
}
func (h *Human) printInfo() {
fmt.Println("My name is", h.Name) //Метод 
}

type MoreInfo struct {
	Height float32
	Weight float32
	Human //композиция
}

func main() {
	// Вложили в переменную MyHuman информацию по человеку с именем Dan
	MyHuman := MoreInfo{
		Height: 180,
		Weight: 62,
		Human: Human{
			Name:  "Dan",
			Hobby: "Programming"}}

	// Выводим данные переменной на экран обратившись к переменной
	MyHuman.printInfo()

	// Выводим данные переменной на экран обратившись к переменной с определенным полем
	fmt.Println(MyHuman.Name, MyHuman.Hobby, MyHuman.Height)
}
