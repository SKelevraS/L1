package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{} //вызов ожидания завершения операций в переменной

	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			 array := sl[i] * sl[i]
			//  sl := i*i
			// fmt.Println(sl)
			fmt.Println(array)
		}(i)
	}

	wg.Wait()
}