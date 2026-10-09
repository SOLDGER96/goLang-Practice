package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/interfaces/note"
	"example.com/interfaces/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

//you can also embed other interfaces in another interface
type outputable interface {
	saver
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {
	//
	printSomething(1)
	printSomething(1.5)
	printSomething("abc")

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

	err = outputData(todo)
	if err != nil { 
		return
	}


	err = outputData(userNote)
	if err != nil {
		return
	}
	
}

// Go any value allowed type
func printSomething(value any){  //interface{} == any
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Integer: ", floatVal)
		return
	}

	strVal, ok := value.(string)
	if ok {
		fmt.Println("Integer: ", strVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }
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

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the Note Failed!")
		return err
	}
	fmt.Println("Saving the note Succeeded!")
	return nil
}

func outputData (data outputable) error {
	data.Display()
	return saveData(data)
}