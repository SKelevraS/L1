package main

import (
	"fmt"
	"sync"
)

type Counter struct { // Создаем структуру, в которой будет находиться наш счетчик
	FinishCount int
	sync.Mutex
}

func main() {
	count := Counter{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ { // Запускаем 100 горутин
		wg.Add(1)
		go func() {
			count.Lock()
			defer wg.Done()
			count.FinishCount++            
			fmt.Println(count.FinishCount) 
			count.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(count.FinishCount) // Выводим общее значение счетчика
}