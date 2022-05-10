package main

import (
	"fmt"
	"time"
)

// Функция записи в основной канал
func write(ch chan int, val int) {
	ch <- val
}

// функция чтения из канал, где записали данные основного канала
func read(ch chan int) {
	v := <-ch
	fmt.Println("Received data with value: ", v) //Показываем визуальное действие
}

func main() {
	var t time.Duration 

	fmt.Print("Set time in seconds: ")
	fmt.Scan(&t) 

	ch := make(chan int)     //Создаем канал, в который мы будем писать и читать
	timer := time.NewTimer(t * time.Second) 

	i := 0
	for {
		select {
		default:
			go write(ch, i) 
			go read(ch)

			time.Sleep(100 * time.Millisecond) //Чтобы не итерировались по числам слишком быстро
			i++      
			
		case <-timer.C: // получаем сигнал по истечению времени 
			fmt.Println("time is up")
			close(ch) // закрываем основной канал по истечению времени
			return

		}
	}
}
