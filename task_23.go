package main

import "fmt"

func main() {
	var index int
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array: ", array)
	fmt.Println("Enter the index of the element to be removed: ")
	fmt.Scan(&index)
	if index >= len(array) {
		fmt.Println("Incorect index")
		return
	}
	copy(array[index:], array[index+1:]) // Копируем array[index+1:] в array[index:]
	array[len(array)-1] = 0
	array = array[:len(array)-1]
	fmt.Println("New array:", array)

}
