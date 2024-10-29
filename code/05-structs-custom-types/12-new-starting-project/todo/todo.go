package todo

import (
	"fmt"
  "os"
  "encoding/json"
)
import "errors"

type Todo struct {
	Text  string `json:"content"`
}

func (todo Todo) Save() error {
  fileName := "todo"

  json, error := json.Marshal(todo)
  
  if error != nil {
    return error
  }

  return os.WriteFile(fileName + ".json", json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Text is required")
	}

	return Todo{text}, nil
}

func (n Todo) Display() {
	fmt.Printf("Your todo text %v\n\n", n.Text)
}
