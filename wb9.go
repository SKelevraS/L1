package main

import "fmt"

func main() {
  m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  squareChannel := make(chan int, 1)
  xChannel := make(chan int, 1)

  go func() {
    for _, i := range m {
      squareChannel <- i
    }
    close(squareChannel)
  }()

  go func() {
    for num := range squareChannel {

      xChannel <- num * 2
    }
    close(xChannel)
  }()

 
    for num := range xChannel {
      fmt.Println(num)
    }

}