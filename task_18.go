package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Counter struct {
	value      int // значение счётчика
	sync.Mutex     // мьютекс,т.к. придётся блокировать горутины иначе может произойти одновременное обращение к одному участку памяти
}

func increment(ctx context.Context, wg *sync.WaitGroup, counter *Counter) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Если контекст завершил работу
			return
		default:
			counter.Mutex.Lock()   // Лочим мьютекс,чтобы доступ был только у одной горутины
			counter.value++        // Инкрементируем счётчик
			counter.Mutex.Unlock() // Анлочим мьютекс
		}
	}
}

func main() {
	var timeOut int
	var counter Counter
	fmt.Println("Lead time: ")
	fmt.Scan(&timeOut)
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(timeOut)*time.Second) // Контекст, живущий timeOut секунд
	defer cancel()
	wg := sync.WaitGroup{}     // Механизм ожидания завершения группы задач
	numCPU := runtime.NumCPU() // Получаем кол-во логических ядер процессора
	wg.Add(numCPU)             // Добавление в группу кол-ва заданий равное ядрам процессора

	for i := 0; i < numCPU; i++ { // Создаём горутины в кол-ве равному числу логических ядер
		go increment(ctx, &wg, &counter) // Инкрементируем счётчик
	}
	wg.Wait() // Метод Wait для блокировкаи основной задачи,чтобы остальные в WaitGroup успели завершиться.
	fmt.Println("Counter value: ", counter.value)
}
