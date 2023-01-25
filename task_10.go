package main

import "fmt"

func main() {
	array := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32) // Мапа,где ключ-температурная метка,значение-элемент массива
	for _, elem := range array {
		mark := int(elem*0.1) * 10 // Выделение температурных меток
		result[mark] = append(result[mark], elem)
	}
	fmt.Println(result)
}
