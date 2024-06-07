package main

import (
	"time"
)

func main() {

	var source ILogger

	logType := "file"
	hasContext := true
	hasEncryption := false
	message := "It is a simple message."
	userId := 1
	sessionId := 2
	encryptionShift := 3

	switch logType {
	case "file":
		source = fileLogger{
			data: map[string]interface{}{
				"message": message,
			},
		}
	default:
		source = consoleLogger{
			data: map[string]interface{}{
				"message": message,
			},
		}
	}

	if hasContext {
		source = contextualLogger{
			logger: source,
		}
		source.setData("userId", userId)
		source.setData("sessionId", sessionId)
		source.setData("timestamp", time.Now())
	}

	if hasEncryption {
		source = encryptionLogger{
			encryptionShift: encryptionShift,
			logger:          source,
		}
	}

	source.log()
}
