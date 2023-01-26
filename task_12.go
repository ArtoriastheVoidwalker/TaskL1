package main

import "fmt"

func RemoveDuplicate(set []string) []string { // Функция для удаления дублекатов
	keys := make(map[string]bool)
	list := []string{}

	for _, elem := range set {
		if _, value := keys[elem]; !value {
			keys[elem] = true
			list = append(list, elem)
		}
	}
	return list
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(RemoveDuplicate(strings)) // Множество не содержит дубликаты элементов
}
