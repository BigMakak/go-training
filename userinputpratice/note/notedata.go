package notedata

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type NoteData struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

func New(title, content string) (*NoteData, error) {
	if title == "" || content == "" {
		return nil, errors.New("title or content is empty")
	}
	return &NoteData{Title: title, Content: content, CreateAt: time.Now()}, nil
}

func (n *NoteData) Display() {
	output := fmt.Sprintf("\n\nNote title: %s \n\nContent: %s \n\nCreatedAt: %s", n.Title, n.Content, n.CreateAt.String())
	fmt.Println(output)
}

func (n *NoteData) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = fmt.Sprintf("%s.%s", strings.ToLower(fileName), "json")

	jsonOutput, err := json.Marshal(n)

	if err != nil {
		return errors.New("error marshalling note data")
	}

	return os.WriteFile(fileName, jsonOutput, 0644)
}
