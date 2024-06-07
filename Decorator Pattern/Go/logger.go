package main

type ILogger interface {
	setData(key string, value interface{}) bool
	log() 
	getData() map[string]interface{}
	delData(key string) bool
}
