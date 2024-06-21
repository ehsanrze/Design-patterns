package main

import (
	"time"
)

func main() {
	fleetPublisher := FleetPublisher{}

	dashboardObserver1 := DashboardObserver{
		id: 1,
	}
	fleetPublisher.registerObserver(&dashboardObserver1)

	dashboardObserver2 := DashboardObserver{
		id: 2,
	}
	fleetPublisher.registerObserver(&dashboardObserver2)

	fleetPublisher.positionUpdated(map[string]interface{}{
		"courierId": 1,
		"lat":       123,
		"lng":       232,
		"timestamp": time.Now().String(),
	})

	fleetPublisher.unregisterObserver(&dashboardObserver1)

	fleetPublisher.courierDisconnected(map[string]interface{}{
		"courierId": 1,
		"lat":       123,
		"lng":       232,
		"timestamp": time.Now().String(),
	})
}
