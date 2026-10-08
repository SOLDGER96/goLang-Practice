package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}



func New(title, content string) (Note, error) {
	if title == "" && content == "" {
		return Note{}, errors.New("Invalid Input")
	}
	return Note{
		title: title,
		content: content,
		createdAt: time.Now(), 
	}, nil
}
func (note Note) Display() {
	fmt.Printf("Your Note titled %s has the following content: \n%s\n", note.title, note.content)
}