package note

import (
	"fmt"
	"time"
)
import "errors"

type Note struct {
	title    string
	content  string
	createAt time.Time
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{title, content, time.Now()}, nil
}

func (n Note) Display() {
	fmt.Printf("You note titled %v has the following content:\n\n%v", n.title, n.content)
}
