package main

import (
	"fmt"
	"os"
)

type fileLogger struct {
	data map[string]interface{}
}

func (l fileLogger) log() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	for key, value := range l.getData() {
		line := fmt.Sprintf("Key: %s, Value: %v\n", key, value)

		if _, err := file.WriteString(line); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Data appended to log.txt successfully")
}

func (l fileLogger) setData(key string, value interface{}) bool {
	l.data[key] = value

	return true
}

func (l fileLogger) delData(key string) bool {
	delete(l.getData(), key)

	return true
}

func (l fileLogger) getData() map[string]interface{} {
	return l.data
}
