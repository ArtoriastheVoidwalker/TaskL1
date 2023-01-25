package main

import (
	"fmt"
	"os"
)

func main() {
	var number, position, bit int64

	fmt.Println(" -------------------------------------------------\n" +
		"| Exemple: number - 65; position - 4; bit - 1;   |\n" +
		"| -----------------------------------------------|\n" +
		"| positions number               : 6 5 4 3 2 1 0 |\n" +
		"| number in binary representation: 1 0 0 0 0 0 1 |\n" +
		"| add bit for position           :     ^         |\n" +
		"| result                         : 1 0 1 0 0 0 1 |\n" +
		" -------------------------------------------------\n")

	fmt.Println("Enter number:")
	fmt.Scan(&number)
	fmt.Println("Choose position bit:")
	fmt.Scan(&position)
	if position < 0 {
		fmt.Println("Position cannot be less than zero.")
		os.Exit(1)
	}
	fmt.Println("Select the bit value 1 or 0:")
	fmt.Scan(&bit)
	if bit != 1|0 {
		fmt.Println("Only binary numbers are allowed.")
		os.Exit(1)
	}
	if bit == 1 {
		number |= (1 << position) // Сдвигает битовое представление числа, представленного первым операндом,
		// влево на определенное количество разрядов, которое задается вторым операндом.
		// |: поразрядная дизъюнкция (операция ИЛИ или поразрядное сложение).
		// Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
	} else {
		number &= ^(1 << position) //  сброс бита (И НЕ). В выражении z = x &^ y каждый бит z равен 0,
		// если соответствующий бит y равен 1. Если бит в y равен 0, то берется значение соответствующего бита из x.
	}
	fmt.Println("Result value:", number)
}