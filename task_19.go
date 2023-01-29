package main

import "fmt"

func main() {
	var reverseStr string
	str := "главрыба"

	for _, v := range str {
		reverseStr = string(v) + reverseStr // Поочерёдно довавляем элементы в начало строки
	}
	fmt.Println(reverseStr)
}
