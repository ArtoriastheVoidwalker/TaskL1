package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var runningTime int // Время жизни программы
	fmt.Println("Program running time:")
	fmt.Scan(&runningTime)
	ch := make(chan int) // Небуферизированный канал для чтения и записи
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(runningTime)*time.Second) // Контекст, сообщающий о завершении
	defer cancel() // Отмена контекста после завершения программы

	go func(ch chan int) { // Горутина для записи последовательных значений в канал с задержкой в Millisecond * 100
		number := 0
		for {
			number++
			ch <- number
			time.Sleep(time.Millisecond * 100)
		}
	}(ch)
	for {
		select {
		case <-ctx.Done(): // Обработка случая, когда контекст завершает работу
			fmt.Println("Timeout")
			return
		default: // Обработка стандартного случая. Чтение данных из канала
			number := <-ch
			fmt.Println(number)
		}
	}
}
