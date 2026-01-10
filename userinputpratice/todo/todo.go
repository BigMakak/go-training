package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (*Todo, error) {
	if content == "" {
		return nil, errors.New("title or content is empty")
	}
	return &Todo{Text: content}, nil
}

func (t *Todo) Display() {
	output := fmt.Sprintf("Todo: %s", t.Text)
	fmt.Println(output)
}

func (t *Todo) Save() error {
	fileName := "todo.json"

	jsonOutput, err := json.Marshal(t)

	if err != nil {
		return errors.New("error marshalling note data")
	}

	return os.WriteFile(fileName, jsonOutput, 0644)
}
