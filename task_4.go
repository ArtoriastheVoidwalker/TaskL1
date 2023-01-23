package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func outputFromChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		_, err := fmt.Fprintln(os.Stdout, val) // Вывод данных из канала
		if err != nil {
			return
		}
	}
}

func main() {
	var numberWorker int                                  // Количество воркеров
	wg := sync.WaitGroup{}                                // Механизм ожидания завершения группы задач
	sigCh := make(chan os.Signal, 1)                      // Буферизированный канал,из которого будем читать отправленные сигналы
	ch := make(chan int)                                  // Небуферизированный канал для данных
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) // signal.Notify регистрирует данный канал для получения уведомлений об указанных сигналах.
	// Сервер корректно завершит работу при получении SIGTERM, или остановит работу, если он получил SIGINT(Сочетанием Ctrl+c)
	fmt.Println("Enter number of workers: ")
	fmt.Scan(&numberWorker) // Получаем количество воркеров из потока ввода
	if numberWorker <= 0 {  // Проверка количества воркеров, при значении меньше нуля данные не будут генерироваться
		fmt.Println("numberWorker must be greater than zero.")
		close(ch) // Закрытие каналов для данных и сигналов
		close(sigCh)
		os.Exit(1) // Завершение работы программы
	}
	for i := 0; i < numberWorker; i++ { // Цикл для чтения данных из канала
		wg.Add(1) // Количество задач,которое WaitGroup должна подождать
		go outputFromChannel(ch, &wg)
	}
	func() {
		for {
			select {
			case <-sigCh: // Обработка случая, если пришел сигнал о завершении работы
				fmt.Println("\nSuccessful completion.")
				return
			default: // Обработка стандартного случая. Запись в канал случайного числа каждые 100 Millisecond
				ch <- rand.Int() // отправляем рандомное значение в канал
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	close(ch) // Закрытие каналов для данных и сигналов
	close(sigCh)
	wg.Wait() // Метод Wait для блокировкаи основной задачи,чтобы остальные в WaitGroup успели завершиться.
}
