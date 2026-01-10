package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/user-input-pratice/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := notedata.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note created and saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your Note: ")

	content := getUserInput("Enter the content of your Note: ")

	return title, content
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
