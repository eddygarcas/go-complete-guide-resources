package note

import (
	"fmt"
	"time"
  "os"
  "encoding/json"
  "strings"
)
import "errors"

type Note struct {
	Title    string
	Content  string
	CreateAt time.Time
}

func (note Note) Save() error {
  fileName := strings.ReplaceAll(note.Title, " ", "_")
  fileName = strings.ToLower(fileName)

  json, error := json.Marshal(note)
  
  if error != nil {
    return error
  }

  return os.WriteFile(fileName + ".json", json, 0644)
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{title, content, time.Now()}, nil
}

func (n Note) Display() {
	fmt.Printf("You note titled %v has the following content:\n\n%v\n\n", n.Title, n.Content)}
