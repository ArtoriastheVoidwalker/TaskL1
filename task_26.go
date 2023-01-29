package main

import (
	"fmt"
	"strings"
)

func unique(currentString string) bool {
	status := true
	mp := make(map[string]int)
	currentString = strings.ToLower(currentString) // Приведение всей строки к нижнему регистру, чтобы обеспечить регистронезависимость
	for _, val := range currentString {            // Подсчёт повторений для каждого символа строки
		if _, ok := mp[string(val)]; ok {
			return false
		} else {
			mp[string(val)]++
		}
	}
	return status
}

func main() {
	var currentString string
	fmt.Println("Enter string: ")
	fmt.Scan(&currentString)

	fmt.Println(unique(currentString))
}
