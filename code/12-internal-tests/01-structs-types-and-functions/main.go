package main

import (
	"example.com/clients/employee"
	"example.com/clients/manager"
)

func main() {

	manager, _ := manager.New("Eduard", "Garcia")
	employee, _ := employee.New("Ramon", "Massana")

	manager.Display(func(element manager.Manager) string { return element.FirstName + " " + element.Surname })
	employee.Display(func(element Element) string { return element.FirstName + " and " + element.Surname })

}
