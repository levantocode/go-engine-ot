package events

type EventDispatcher struct {
	listeners map[string][]func(Event) // map of event name & its functions being dispatched
}

func (eventDispatcher *EventDispatcher) Subscribe(eventName string, callback func(Event)) {
	eventDispatcher.listeners[eventName] = append(eventDispatcher.listeners[eventName], callback)
}

func (eventDispatcher *EventDispatcher) Dispatch(event Event) {
	if callbacks, exists := eventDispatcher.listeners[event.Name]; exists {
		for _, callback := range callbacks {
			callback(event)
		}
	}
}
