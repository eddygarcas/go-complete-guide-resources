package manager

import "fmt"

type Manager struct {
	FirstName string
	Surname   string
	role      string
}

func (e Manager) Display(formatName func(element Manager) string) {
	// fmt.Printf("Your full name is %v %v\n", e.FirstName, e.Surname)
	fmt.Printf("Manager name: %v\n", formatName(e))
}

func New(firstName, surname string) (Manager, error) {
	return Manager{
		FirstName: firstName,
		Surname:   surname,
	}, nil
}
