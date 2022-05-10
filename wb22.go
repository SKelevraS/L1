package main

import "fmt" //math/big

func main() {
	var A int
	_, aErr := fmt.Scan(&A)
	checkError(aErr)

	var B int
	_, bErr := fmt.Scan(&B)
	checkError(bErr)

	C := A * B
	fmt.Println("Multiplication:", C)
	C = A + B
	fmt.Println("Sum:", C)
	C = A - B
	fmt.Println("Subtraction:", C)
	C = A % B
	fmt.Println("Remainder:", C)
	C = A / B
	fmt.Println("Division:", C)
	var D float64
	D = float64(A) / float64(B)
	fmt.Println("Division:", D)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}