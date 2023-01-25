package main

import (
	"fmt"
)

func removeDuplicate(set []string) []string { // Функция для удаления дублекатов(Если не удалять дублиаты,
	// то одиннаковые значения одного множества сойдут за уникальные для обоих)
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
	cjOrder := []string{"number 9"}                                   // Заказ Карла Джонсона
	smokesOrder := []string{"number 9", "number 9", "number 9 large", // Заказ Мелвина Харриса
		"number 6 with extra dip",
		"number 7",
		"number 45",
		"number 45 with cheese",
		"large soda"}
	counter := make(map[string]int)
	intersection := make([]string, 0)
	for _, elem := range removeDuplicate(cjOrder) {
		counter[elem] += 1
	}

	for _, elem := range removeDuplicate(smokesOrder) {
		counter[elem] += 1
	}

	for key := range counter {
		if counter[key] > 1 {
			intersection = append(intersection, key)
		}
	}
	fmt.Println(intersection)
}
