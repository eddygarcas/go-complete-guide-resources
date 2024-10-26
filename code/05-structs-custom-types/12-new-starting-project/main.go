package main

import (
	"example.com/notes-app/note"
	"fmt"
)

func main() {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
}

func getUserInput(prompt string) string {
	var userInput string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return ""
	}

	if userInput == "" {
		fmt.Println("You must enter a value")
		return getUserInput(prompt)
	}

	return userInput
}
