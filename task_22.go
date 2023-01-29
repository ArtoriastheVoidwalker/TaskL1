package main

import (
	"fmt"
	"math/big"
)

func getSum(a int64, b int64) big.Int { // Функция для получения суммы
	var sum big.Int
	x := big.NewInt(a) // Выделение памяти под новое слогаемое
	y := big.NewInt(b)
	sum.Add(x, y)
	return sum
}

func getDifference(a int64, b int64) big.Int { // Функция для получения разницы
	var difference big.Int
	x := big.NewInt(a)
	y := big.NewInt(b)
	difference.Sub(x, y)
	return difference
}

func getProduct(a int64, b int64) big.Int { // Функция для получения произведения
	var product big.Int
	x := big.NewInt(a)
	y := big.NewInt(b)
	product.Mul(x, y)
	return product
}

func getQuotient(a int64, b int64) big.Int { // Функция для получения частного
	var quotient big.Int
	x := big.NewInt(a)
	y := big.NewInt(b)
	quotient.Div(x, y)
	return quotient
}

func main() {
	var a, b int64 = 53443553443553443, 32112332112332112
	fmt.Println(getSum(a, b))
	fmt.Println(getDifference(a, b))
	fmt.Println(getProduct(a, b))
	fmt.Println(getQuotient(a, b))

}
