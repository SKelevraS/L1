package main

import "fmt"

func main() {

	var word string
	_, err := fmt.Scan(&word)
	checkError(err)

	var ReverseWord string

	for _, i := range word {
		ReverseWord = string(i) + ReverseWord
		fmt.Println(ReverseWord)

	}
	fmt.Println(ReverseWord)

}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
