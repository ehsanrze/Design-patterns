package main

import "fmt"

type DashboardObserver struct {
	id int
}

func (do *DashboardObserver) update(data map[string]interface{}) {
	fmt.Printf(
		"ObserverId: %d, Event: %s \n CourierId: %d \n Position: %d,%d \n Timestamp: %s \n\n\n",
		do.id,
		data["type"],
		data["courierId"],
		data["lat"],
		data["lng"],
		data["timestamp"])
}
