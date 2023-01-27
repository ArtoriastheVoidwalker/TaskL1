package main

import (
	"fmt"
)

func main() {
	origin := []int{5, 6, 7, -4, 5, 6, 8, -3, 5, 8}
	quickSort(origin, 0, len(origin))
	fmt.Println("quickSort: ", origin)
}

func quickSort(array []int, left, right int) {
	if right > left {
		mid := partition(array, left, right) // Получение центрального элемента массива
		quickSort(array, left, mid)          // рекурсивно перебираем значения в диапазоне от нулевого элемента до середины и
		quickSort(array, mid+1, right)       // от середины до последнего
	}
}

func partition(array []int, left, right int) (index int) { // Функция отвечает за поиск индекса "центрального элемента"
	// и распределение элементов по разные стороны от него
	elem := array[left]
	index = left
	for j := index + 1; j < right; j++ {
		if array[j] < elem {
			index++
			array[index], array[j] = array[j], array[index]
		}
	}
	array[index], array[left] = array[left], array[index]
	return index
}
