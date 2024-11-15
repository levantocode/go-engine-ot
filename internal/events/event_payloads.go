package events

type ServerStatusChangePayload struct {
	NewStatus ServerStatus
}

type ServerStatus int

const (
	ServerStatusStarting ServerStatus = iota
	ServerStatusPlaying
	ServerStatusRestarting
	ServerStatusShuttingDown
)
