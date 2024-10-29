package main

import (
	"example.com/notes-app/note"
	"fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
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
