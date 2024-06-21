package main

type Publisher interface {
	registerObserver(observer Observer)
	unregisterObserver(observer Observer)
	notifyObservers()
}
