package main

import (
	"fmt"
)

func binarySearch(a []int, target int) (result int) {
	mid := len(a) / 2 // находим "центральный элемент массива"
	switch {
	case len(a) == 0: // Если пришел пустой массив
		result = -1
	case a[mid] > target: // Если искомый элемент меньше центрального элемента
		result = binarySearch(a[:mid], target) // Рекурсивно вызываем функию и сдвигаем верхнюю границу поиска до mid
	case a[mid] < target: // Если искомый элемент больше центрального элемента
		result = binarySearch(a[mid+1:], target) // Рекурсивно вызываем функию и сдвигаем нижнюю границу поиска до mid
		if result >= 0 {                         // Если искомый элемент между границами
			result += mid + 1
		}
	default:
		result = mid
	}
	return
}

func main() {
	array := []int{1, 2, 4, 7, 9, 18, 21, 34, 35, 40}
	fmt.Println(binarySearch(array, 21))

}
