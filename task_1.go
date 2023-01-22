package main

import "fmt"

const (
	low    string = "low"
	medium string = "medium"
	hard   string = "hard"
)

type Human struct {
	name    string
	age     int
	cuntry  string
	address string
	phone   string
}

type Action struct {
	Human
	task     string
	priority string
}

func (h *Human) info() string {
	return fmt.Sprintf("Name = %s, \n Age = %d,\n Cuntry = %s,\n Address = %s,\n Phone = %s",
		h.name, h.age, h.cuntry, h.address, h.phone)
}

func (h *Human) currentTask(task string) {
	fmt.Println("My current task: ", task)
}

func main() {
	action := Action{
		Human: Human{
			name:    "Robert «Bob» Paulson",
			age:     35,
			cuntry:  "US",
			address: "Paper street",
			phone:   "+12345678901",
		},
		task:     "Participation in the project «Destruction»",
		priority: hard,
	}
	// Встраивание Human в Action делает возможным применения медодов Human к объекту Action
	fmt.Println(action.info())
	action.currentTask(action.task)
}
