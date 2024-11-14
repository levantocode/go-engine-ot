package events

type Event struct {
	Name    string      // Event type identifier (use string as enum)
	Payload interface{} // Event-specific data
}
