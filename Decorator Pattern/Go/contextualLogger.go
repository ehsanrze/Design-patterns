package main

type contextualLogger struct {
	logger ILogger
}

func (l contextualLogger) log() {

	l.logger.log()
}

func (l contextualLogger) setData(key string, value interface{}) bool {
	return l.logger.setData(key, value)
}

func (l contextualLogger) getData() map[string]interface{} {
	return l.logger.getData()
}

func (l contextualLogger) delData(key string) bool {
	return l.logger.delData(key)
}
