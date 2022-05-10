package main

import (
	"context"
	"time"
)

func main() {
	// Обычное завершение горутины
	go func() {
	}()

	// Завершение горутины через return
	go func() {
		return
	}()


	  // Завершение горутиы через select, закрытие канала и return
		ch := make(chan struct{})
		close(ch)
	go func() {
		select {
		case <-ch:
			return
		}
	}()


	// Завершение горутиы через select, таймер и return
  timer := time.NewTimer(time.Second) //Лучше объявить в самом main
	go func() {
		select {
		case <-timer.C:
			return
		}
	}()

	// Завершение горутиы через select, context и return
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
	go func() {
		select {
		case <-ctx.Done():
			return
		}
	}()

}