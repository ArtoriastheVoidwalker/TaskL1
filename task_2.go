package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // Механизм ожидания завершения группы задач,в который циклично будем добавлять их
	for _, elem := range array {
		wg.Add(1) // Количество задач,которое WaitGroup должна подождать
		go func(elem int) {
			squared := elem * elem
			fmt.Println(elem, " squared ", squared)
			defer wg.Done() // По завершению работы функции к wg применяется метод Done(),для завершения задач
		}(elem)
		wg.Wait() // Метод Wait для блокировкаи основной задачи,чтобы остальные в WaitGroup успели завершиться.
		// Как только счётчик wg будет 0 тогда основная задача разблокируется.
	}
}
