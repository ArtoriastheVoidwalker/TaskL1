package main

import (
	"fmt"
	"strings"
)

/*
	Строки "по капотом" слайс байт или рун, в зависимости от того что за символы хранит, если аски символы, то
одного байта будет достаточно на символ,если руна 2-4, в зависимости от символа. Из этого следует что если брать фрагмент
строки то может возникнуть ошибка т.к. итерация идёт по байтам, а не по символам.
	Нет необходимости выводить justString в глобальную область видимости.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)[:100]
	m := make([]rune, 100)
	copy(m, runes)
	justString = string(m)
	fmt.Println(justString)
}

func createHugeString(len int) string {
	var sb strings.Builder
	for i := 0; i < len; i++ {
		sb.WriteString("★")
	}
	return sb.String()
}

func main() {
	someFunc()
}
