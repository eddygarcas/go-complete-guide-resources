package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes-app/note"
	"example.com/notes-app/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputter interface {
	saver
	displayer
}

func main() {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	todoText := getUserInput("Todo text:")

	printSomething(1)
	printSomething("Hello")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(todo)
	outputData(userNote)
}

func printSomething(value interface{}) {
	typeVal, ok := value.(int)
	if ok {
		fmt.Println("Int value:", typeVal+1)
	} else {
		fmt.Println("Unknown value:", value)
	}
	//switch value.(type) {
	//case string:
	//	fmt.Println("String value:", value)
	//case int:
	//	fmt.Println("Int value:", value)
	//default:
	//	fmt.Println("Unknown value:", value)
	//}
}

func outputData(data outputter) error {
	data.Display()
	err := saveData(data)
	if err != nil {
		fmt.Println("Error saving data")
		return err
	}
	fmt.Println("Data saved")
	return nil
}

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println("Error saving data")
		return err
	}
	fmt.Println("Data saved")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")
	userInput = strings.TrimSuffix(userInput, "\r")

	if userInput == "" {
		fmt.Println("You must enter a value")
		return getUserInput(prompt)
	}

	return userInput
}
