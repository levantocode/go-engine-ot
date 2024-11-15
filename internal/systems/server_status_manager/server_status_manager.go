package server_status_manager

import (
	"fmt"

	"github.com/levantocode/go-engine-ot/internal/events"
)

var currentServerStatus events.ServerStatus

func Init(dispatcher *events.EventDispatcher) {
	dispatcher.Subscribe(events.ServerStatusChangeEvent, func(e events.Event) {
		if payload, ok := e.Payload.(events.ServerStatusChangePayload); ok {
			updateCurrentServerStatus(payload.NewStatus)
		}
	})
}

func updateCurrentServerStatus(newStatus events.ServerStatus) {
	if currentServerStatus == events.ServerStatusShuttingDown {
		return // No stopping Shut Down
	}

	if currentServerStatus == events.ServerStatusPlaying && newStatus == events.ServerStatusStarting {
		fmt.Print("Invalid Server State Change Attempt: tried to Change Server Status from Playing to Starting.\n")
		return
	}

	if currentServerStatus == newStatus {
		fmt.Printf("Tried to Change Server Status to: %v but it was already in it.\n", newStatus)
		return
	}

	currentServerStatus = newStatus
	fmt.Printf("Changing Server Status to: %v\n", newStatus)
}
