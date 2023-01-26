package main

import "fmt"

func getType(elem interface{}) {
	switch expr := elem.(type) {
	case int:
		fmt.Println("Type ", expr, " is int")
	case string:
		fmt.Println("Type ", expr, " is string")
	case bool:
		fmt.Println("Type ", expr, " is bool")
	case chan int:
		fmt.Println("Type ", expr, " is channel")
	}
}

func main() {
	var interfaceInt interface{} = 1
	var interfaceString interface{} = "one"
	var interfaceBool interface{} = true
	var interfaseChannel interface{} = make(chan int)

	getType(interfaceInt)
	getType(interfaceString)
	getType(interfaceBool)
	getType(interfaseChannel)
}
