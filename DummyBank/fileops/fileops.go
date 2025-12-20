package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(filePath string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return defaultValue, errors.New("unable to read balance.txt, file does not exist")
	}
	amount, errorParse := strconv.ParseFloat(string(data), 64)

	if errorParse != nil {
		return defaultValue, errors.New("unable to parse balance from file")
	}

	return amount, nil
}

func WriteFloatToFile(filePath string, amount float64) error {
	balanceText := fmt.Sprint(amount)
	err := os.WriteFile(filePath, []byte(balanceText), 0644)

	if err != nil {
		fmt.Println("Error writing balance to file")
		return err
	}

	return nil
}
