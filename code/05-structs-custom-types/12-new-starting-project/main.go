package main

import (
  "fmt"
  "example.com/12-new-starting-project/note"
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
  fmt.Scanln(&userInput)

  if userInput == "" {
    fmt.Println("You must enter a value")
    return getUserInput(prompt)
  }

  return userInput
}
