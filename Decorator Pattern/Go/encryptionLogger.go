package main

import (
	"encoding/json"
)

type encryptionLogger struct {
	encryptionShift int
	logger          ILogger
}

func (l encryptionLogger) log() {

	data := l.logger.getData()

	for key, value := range data {

		encryptKey, err := l.encrypt(key, l.encryptionShift)

		if err != nil {
			return
		}

		encryptValue, err := l.encrypt(value, l.encryptionShift)
		if err != nil {
			return
		}

		l.logger.delData(key)

		l.logger.setData(encryptKey, encryptValue)

	}

	l.logger.log()
}

func (l encryptionLogger) encrypt(data interface{}, shift int) (string, error) {

	plaintext, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	var encrypted string
	for _, char := range string(plaintext) {
		switch {
		case char >= 'a' && char <= 'z':
			encrypted += string('a' + (char-'a'+rune(shift))%26)
		case char >= 'A' && char <= 'Z':
			encrypted += string('A' + (char-'A'+rune(shift))%26)
		case char >= '0' && char <= '9':
			encrypted += string('0' + (char-'0'+rune(shift))%10)
		default:
			encrypted += string(char)
		}
	}
	return encrypted, nil
}

func (l encryptionLogger) setData(key string, value interface{}) bool {
	return l.logger.setData(key, value)
}

func (l encryptionLogger) getData() map[string]interface{} {
	return l.logger.getData()
}

func (l encryptionLogger) delData(key string) bool {
	return l.logger.delData(key)
}
