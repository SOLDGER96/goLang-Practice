package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/interfaces/note"
	"example.com/interfaces/todo"
)

func main() {
	//
	title, content := getNoteData()

	todoText := getUserInput("Todo Text: ")
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
		fmt.Println("Saving the Note Failed!")
		return
	}
	fmt.Println("Saving the todo Succeeded!")


	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the Note Failed!")
		return
	}
	fmt.Println("Saving the note Succeeded!")
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ",prompt)
	

	// Scan can only be used for inputs w/o a space in between 
	// for longer inputs we have different approach
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil { 
		return ""
	}
	// Cleaning text

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")  // Line break for Windows
	

	return text
}
func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	
	return title, content
}
