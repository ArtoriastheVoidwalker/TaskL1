package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func newSleep(value int) {
	select {
	case <-time.After(time.Duration(value) * time.Second): // Канал возвращающий значение текущее времени и завершающий работу
		return
	}
}

func main() {
	var sec int
	fmt.Println("Enter time: ")
	fmt.Scan(&sec)
	fmt.Println("Start sleep:", time.Now()) // Получает текущее время с начала выполнения функции
	newSleep(sec)
	fmt.Println("Stop sleep:", time.Now()) // и после завершения её работы,чтобы убедиться в её работоспособности
}
