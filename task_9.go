package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup // Механизм ожидания завершения группы задач
	array := [5]int{2, 4, 6, 8, 10}
	channelX := make(chan int)  // Канал для записи элемнтов из массива
	channel2X := make(chan int) // Канал для записи элемнтов *2 из массива
	wg.Add(1)

	go func() { // Запись элементов из массив в канал channelX
		for _, elem := range array {
			channelX <- elem
		}
		close(channelX)
	}()

	go func() { // Запись элементов из канала channelX в channel2X
		for elem := range channelX {
			channel2X <- elem * 2
		}
		close(channel2X)
	}()

	go func() { // Вывод элементов из канала channel2X в stdout
		defer wg.Done()
		for elem := range channel2X {
			_, err := fmt.Fprintln(os.Stdout, elem)
			if err != nil {
				return
			}
		}
	}()

	wg.Wait() // Метод Wait для блокировкаи основной задачи,чтобы остальные в WaitGroup успели завершиться.
}
