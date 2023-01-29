package main

import "fmt"

type Point struct { // Структура точки с коардинатами относительно х и у
	x int
	y int
}

func newPoint(x int, y int) *Point { // Конструктор
	point := Point{x, y}
	return &point
}

func (p *Point) calculateDistance() int { // Метод расчёта расстояния
	if p.x == p.y {
		return 0
	} else if p.x > p.y {
		return p.x - p.y
	} else {
		return p.y - p.x
	}
}

func main() {
	var x, y int
	fmt.Println("Enter x: ")
	fmt.Scan(&x)
	fmt.Println("Enter y: ")
	fmt.Scan(&y)
	point := newPoint(x, y)
	count := point.calculateDistance()
	fmt.Println(count)
}
