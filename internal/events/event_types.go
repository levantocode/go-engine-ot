package events

type Event struct {
	Name    string      // Event type identifier (use string)
	Payload interface{} // Event-specific data
}

type EventType int

const (
	ServerStatusChangeEvent = "server_status_change"
)
