package employee

import (
	"example.com/clients/manager"
	"fmt"
)

type Employee struct {
	FirstName string
	Surname   string
	Manager   manager.Manager
}

func (e Employee) Display(formatName func(element Employee) string) {
	//fmt.Printf("Your full name is %v %v\n", e.FirstName, e.Surname)
	fmt.Printf("Employee name: %v", formatName(e))
}

func (e Employee) Format() {
	e.Display(func(element Employee) string { return element.FirstName + " and " + element.Surname })
}

func New(firstName, surName string) (Employee, error) {
	return Employee{
		FirstName: firstName,
		Surname:   surName,
	}, nil
}
