package main

import (
	"fmt"
	"sync"
)

type Summ struct {
  //композиция метода блокировки замка
  sync.Mutex
  //композиция типа
  someV int
}

func main() {
  var wg sync.WaitGroup

  arr := []int{2, 4, 6, 8, 10}
  wg.Add(len(arr))

  var sum Summ
  for _, e := range arr {
    go func(e int, sum *Summ) {
      sum.Lock()
      sum.someV += e * e
      sum.Unlock()
      defer wg.Done()
    }(e, &sum)
  }
  wg.Wait()
  fmt.Println(sum.someV)
}