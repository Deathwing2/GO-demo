package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Цикл номер:", i)
		time.Sleep(2 * time.Second) // Ожидание 2 секунды
	}
}
