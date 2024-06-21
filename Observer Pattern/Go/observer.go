package main

type Observer interface {
	update(value map[string]interface{})
}
