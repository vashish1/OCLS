package hub

import (
	"sync"

	"github.com/vashish1/OnlineClassPortal/pkg/models"
)

var HubConstruct = models.Hub{
	Broadcast:  make(chan models.WsMessage),
	Register:   make(chan models.Subscription),
	UnRegister: make(chan models.Subscription),
	Users: models.WsUsers{
		Users: make(map[string]map[*models.Connection]bool),
		Mutex: &sync.RWMutex{},
	},
}
