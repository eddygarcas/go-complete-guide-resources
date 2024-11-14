package main

import (
	"example.com/clients/employee"
	"example.com/clients/manager"
)

// This interfice will be implemented by the manager and employee structs
// this way we can call printSomething with a manager or employee
type Person interface {
	Format()
}

func main() {

	manager, _ := manager.New("Eduard", "Garcia")
	employee, _ := employee.New("Ramon", "Massana")

	manager.Format()
	employee.Format()

	printSomething(&manager)
	printSomething(&employee)
}

// Create a generic function where T could be a manager and employee
// and call the Format method on the parameter
func printSomething[T Person](x *T) {
	(*x).Format()
}
