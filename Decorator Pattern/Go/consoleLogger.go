package main

import (
	"fmt"
)

type consoleLogger struct {
	data map[string]interface{}
}

func (l consoleLogger) log() {
	for key, value := range l.getData() {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}

func (l consoleLogger) setData(key string, value interface{}) bool {
	l.data[key] = value

	return true
}

func (l consoleLogger) delData(key string) bool {
	delete(l.getData(), key)

	return true
}

func (l consoleLogger) getData() map[string]interface{} {
	return l.data
}
