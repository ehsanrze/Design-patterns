package main

type FleetPublisher struct {
	observers []Observer
}

func (fp *FleetPublisher) registerObserver(observer Observer) {
	fp.observers = append(fp.observers, observer)
}

func (fp *FleetPublisher) unregisterObserver(observer Observer) bool {
	for index, obs := range fp.observers {

		if obs == observer {
			fp.observers = append(fp.observers[:index], fp.observers[index+1:]...)

			break
		}
	}

	return true
}

func (fp *FleetPublisher) notifyObservers(data map[string]interface{}) {
	for _, observer := range fp.observers {
		observer.update(data)
	}
}

func (fp *FleetPublisher) positionUpdated(data map[string]interface{}) {
	data["type"] = "position_updated"
	fp.notifyObservers(data)
}

func (fp *FleetPublisher) courierDisconnected(data map[string]interface{}) {
	data["type"] = "courier_disconnected"
	fp.notifyObservers(data)
}
