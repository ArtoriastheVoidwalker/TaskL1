package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	setString := strings.Split(str, " ") // Разбиение строки по пробелу

	for elem := 0; elem < len(setString)/2; elem++ { // Меняем местами первый и последний элементы,
		// пока не обошли половину элементов. Данном случае достаточно одной итерации
		lastElem := len(setString) - 1 - elem // Индекс последнего элемента строки
		setString[elem], setString[lastElem] = setString[lastElem], setString[elem]
	}
	fmt.Println(strings.Join(setString, " ")) // "Склеиваем" строку заново
}
