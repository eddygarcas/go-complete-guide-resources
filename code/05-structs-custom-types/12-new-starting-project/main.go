package main

import (
	"example.com/notes-app/note"
  "example.com/notes-app/todo"
	"fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
  todoText := getUserInput("Todo text:")

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

  todo.Display()
  err = todo.Save()

  if err != nil {
    fmt.Println("Error saving todo")
    return
  }
  fmt.Println("Todo saved")

	userNote.Display()
  err = userNote.Save()
  if err != nil {
    fmt.Println("Error saving note")
    return
  }
  fmt.Println("Note saved")
}


func getUserInput(prompt string) string {
	fmt.Printf("%v ",prompt)
  
  userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  userInput = strings.TrimSuffix(userInput,"\n")
  userInput = strings.TrimSuffix(userInput,"\r")
  
  if userInput == "" {
		fmt.Println("You must enter a value")
		return getUserInput(prompt)
	}

	return userInput
}
