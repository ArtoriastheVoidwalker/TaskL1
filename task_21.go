package main

import (
	"fmt"
)

//Паттерн Adapter относится к структурным паттернам уровня класса.
//Его смысл в решении конфликта между классами из-за несовместимости интерфейсов.

//Требуется для реализации:
//- Интерфейс Target, описывающий целевой интерфейс (тот интерфейс с которым система хотела бы работать).
//- Класс Adaptee, который наша система должна адаптировать под себя.
//- Класс Adapter, адаптер реализующий целевой интерфейс.

//В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс.
//Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс.
//Также в языке Go за место общепринятого наследования используется агрегирование и встраивание.

// В качестве примера приведена логика работы оросительной системы. Шланги, подключеные к ней,
// сами по себе не будут поливать растения, полив происходит благодря запуску оросительной системы.

//irrigationSystem — целевой интерфейс
//WaterHose является адаптируемым
//WaterAdapter — это адаптер, реализующий интерфейс освещения.

type WaterHose struct { // Класс WaterHose, который нужно адаптировать к интерфейсу irrigationSystem (Adaptee)
	status bool
}

func (waterHose *WaterHose) GetStatus() bool { // Метот для получения статуса работы системы
	return waterHose.status
}

func (waterHose *WaterHose) StatusOn(On bool) { // Метот для установки статуса системы
	waterHose.status = On
}

type irrigationSystem interface { // Интерфейс оросительной системы — целевой интерфейс
	Watering()
}

type WaterAdapter struct { // Класс адаптера для полива
	waterHose WaterHose
}

func (WaterAdapter *WaterAdapter) Watering() { // Реализация адаптера интерфейса irrigationSystem
	waterHose := WaterAdapter.waterHose

	if waterHose.GetStatus() == false {
		fmt.Println("Status: Irrigation system is off. Startup in progress.")
		waterHose.StatusOn(true)
	} else {
		fmt.Println("Status: Irrigation system is on. Process termination.")
		waterHose.StatusOn(false)
	}
}

func main() {
	var system irrigationSystem = &WaterAdapter{
		waterHose: WaterHose{false},
	}
	system.Watering()
}