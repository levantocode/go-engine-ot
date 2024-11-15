package main

import (
	"github.com/levantocode/go-engine-ot/internal/events"
	"github.com/levantocode/go-engine-ot/internal/systems/server_status_manager"
)

func main() {
	eventDispatcher := events.NewEventDispatcher()
	server_status_manager.Init(eventDispatcher)
}
