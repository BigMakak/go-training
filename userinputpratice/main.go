package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/user-input-pratice/note"
	"example.com/user-input-pratice/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {

	title, content := getNoteData()

	userNote, err := notedata.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todoNote, err := todo.New(getUserInput("Enter a todo item: "))

	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	err = outputData(todoNote)

	if err != nil {
		return
	}

	outputData(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your Note: ")

	content := getUserInput("Enter the content of your Note: ")

	return title, content
}

func outputData(data outputable) error {
	data.Display()
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}

	return nil
}

func getUserInput(question string) string {
	fmt.Printf(" %s", question)

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Trim newline and carriage return characters from the user input
	input = strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")

	return input
}
