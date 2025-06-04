package ai

import (
	"fmt"
	"os"
)

func SetConfig(aiKey string) {
	if aiKey == "" {
		fmt.Println("Cannot set empty AI configuration.")
		return
	}

	file, err := os.OpenFile("key.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(aiKey)
	if err != nil {
		fmt.Println("Error writing to config file:", err)
		return
	}

	fmt.Println("AI configuration set successfully.")
}

func GetConfig() (string, error) {
	file, err := os.Open("key.txt")
	if err != nil {
		return "", fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var aiKey string
	_, err = fmt.Fscanf(file, "%s", &aiKey)
	if err != nil {
		return "", fmt.Errorf("error reading config file: %w", err)
	}

	return aiKey, nil
}