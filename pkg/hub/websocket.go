package hub

import (
	"sync"
	"time"
)

var HubConstruct = Hub{
	Broadcast:  make(chan WsMessage),
	Register:   make(chan Subscription),
	UnRegister: make(chan Subscription),
	Users: WsUsers{
		Users: make(map[string]map[*Connection]bool),
		Mutex: &sync.RWMutex{},
	},
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 30 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 120 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)
