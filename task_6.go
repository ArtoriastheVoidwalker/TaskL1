package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("\n-------------------------\nStopping goroutine execution with WaitGroup\n-------------------------\n")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go waitGroupeStop(&wg)
	wg.Wait()
	fmt.Println("\n-------------------------\nWaitGroup stop\n-------------------------\n")

	fmt.Println("Stopping goroutine execution with WithCancel\n-------------------------\n")
	ctxWithCancel, cancelWithCancel := context.WithCancel(context.Background())
	defer cancelWithCancel()
	go contextWithCancel(ctxWithCancel)
	time.Sleep(time.Microsecond * 35)
	cancelWithCancel()
	fmt.Println("\n-------------------------\nWithCancel stop\n-------------------------\n")

	fmt.Println("Stopping goroutine execution with WithTimeout\n-------------------------\n")
	ctxcontextWithTimeout, cancelcontextWithTimeout := context.WithTimeout(context.Background(), time.Microsecond*25)
	defer cancelcontextWithTimeout()
	go contextWithTimeout(ctxcontextWithTimeout)
	time.Sleep(time.Microsecond * 35)
	fmt.Println("\n-------------------------\nWithTimeout stop\n-------------------------\n")

	fmt.Println("Stopping goroutine execution with WithDeadline\n-------------------------\n")
	ctxWithDeadline, cancelWithDeadline := context.WithDeadline(context.Background(), time.Now().Add(time.Microsecond*25))
	wgWithDeadline := sync.WaitGroup{}
	wgWithDeadline.Add(1)
	defer cancelWithDeadline()
	go contextWithDeadline(ctxWithDeadline, &wgWithDeadline)
	wgWithDeadline.Wait()
	fmt.Println("\n-------------------------\nWithDeadline stop\n-------------------------\n")

}

func waitGroupeStop(wg *sync.WaitGroup) { // Механизм ожидания завершения группы задач. По завершению полезной работы
	fmt.Println("Useful work WaitGroup") // wg.Done() завершит выполнение всех задач в WaitGroup
	time.Sleep(time.Microsecond * 35)
	defer wg.Done()
}

func contextWithCancel(ctx context.Context) { // WithCancel возвращает контекст и функцию, что может его отменить.
	for { // По истенению 35 млсек произойдёт обработка случая, когда контекст завершает работу
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Useful work WithCancel")
			time.Sleep(time.Microsecond * 35)
		}
	}
}

func contextWithTimeout(ctx context.Context) { // WithTimeout вернёт контекст и завершится через указанное время
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Useful work WithTimeout")
		}
	}
}

func contextWithDeadline(ctx context.Context, wg *sync.WaitGroup) { // WithDeadline вернёт контекст и временную метку его завершения
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Useful work WithDeadline")
		}
	}
}
