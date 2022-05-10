package main

import (
	"fmt"
	"time"
)

func MySleep(t time.Duration) {
	 <-time.After(t)
}

func main() {
	fmt.Println("Sleep")

	MySleep(time.Second)

	fmt.Println("two sleep")
}
