package main

import (
	"fmt"
	"sync"
)

func main() {
	var amount int        // Количество значений в мапе
	var wg sync.WaitGroup // Механизм ожидания завершения группы задач
	var mx sync.Mutex     // Механизм синхронизации потоков. Разрешает одному потоку выполнять несколько операций одновременно
	newMap := make(map[int]int)
	fmt.Println("Write amount elements in map:")
	fmt.Scan(&amount)

	for elem := 0; elem <= amount; elem++ {
		wg.Add(1) // Количество задач,которое WaitGroup должна подождать
		go func(elem int) {
			defer wg.Done() // По завершению работы функции к wg применяется метод Done(),для завершения задач
			mx.Lock()       // Лочим мьютекс,чтобы доступ был только у одной горутины
			newMap[elem] = elem
			mx.Unlock() // Анлочим мьютекс
		}(elem)
		wg.Wait() // Метод Wait для блокировкаи основной задачи,чтобы остальные в WaitGroup успели завершиться.
	}
	fmt.Println(newMap)
}
