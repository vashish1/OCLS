package model

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/fasthttp/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 30 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 120 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var HubConstruct = hub{
	broadcast:  make(chan wsMessage),
	register:   make(chan subscription),
	unRegister: make(chan subscription),
	users: wsUsers{
		users: make(map[string]map[*connection]bool),
		mutex: &sync.RWMutex{},
	},
}

func (h *hub) RegisterWS(ws *websocket.Conn, email string) {
	c := &connection{send: make(chan []byte, 256), ws: ws}

	s := subscription{conn: c, user: email}
	HubConstruct.register <- s
	log.Infoln("user", email, "Connected")
	go s.readPump(email)
	go s.writePump()

}

func (h *hub) Run() {
	Upgrader.CheckOrigin = func(r *http.Request) bool {
		host := r.Header.Get("Origin")

		for _, allowdOrigin := range values.Config.CorsAllowedOrigins {
			if host == allowdOrigin {
				return true
			}
		}
		log.Debugln("Host is not allowed:", host)

		return false
	}

	for {
		select {
		case s := <-h.register:
			h.users.mutex.Lock()

			if _, exists := h.users.users[s.user]; !exists {
				connections := make(map[*connection]bool)

				h.users.users[s.user] = connections
			}

			h.users.users[s.user][s.conn] = true

			log.Infoln(s.user, "registered")
			go broadcastOnlineStatusToAllUserRoom(s.user, true)
			h.users.mutex.Unlock()

		case s := <-h.unRegister:
			h.users.mutex.Lock()
			connections, exists := h.users.users[s.user]
			if exists {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.users.users, s.user)
						go broadcastOnlineStatusToAllUserRoom(s.user, false)
						log.Infoln(s.user, "offline")
					}
					log.Infoln(s.user, "subscription removed")
				}
			}
			h.users.mutex.Unlock()

		case m := <-h.broadcast:
			h.users.mutex.RLock()
			connections := h.users.users[m.user]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						h.users.mutex.Lock()
						delete(h.users.users, m.user)
						h.users.mutex.Unlock()
						go broadcastOnlineStatusToAllUserRoom(m.user, false)
					}
				}
			}
			h.users.mutex.RUnlock()
		}
	}
}

func (h *hub) sendMessage(msg []byte, user string) {
	m := wsMessage{msg, user}
	HubConstruct.broadcast <- m
}

// WritePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	if err := c.ws.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		return err
	}

	return c.ws.WriteMessage(mt, payload)
}

// ReadPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump(user string) {
	c := s.conn

	defer func() {
		HubConstruct.unRegister <- s
		c.ws.Close()
	}()

	c.ws.SetReadDeadline(time.Now().Add(pongWait))

	c.ws.SetPongHandler(
		func(string) error {
			return c.ws.SetReadDeadline(time.Now().Add(pongWait))
		})

	for {
		var err error
		var msg messageBytes
		_, msg, err = c.ws.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Errorln("error: ", err)
			}

			return
		}

		data := struct {
			MsgType    string `json:"msgType"`
			User       string `json:"userID"`
			SearchText string `json:"searchText"`
		}{}

		err = json.Unmarshal(msg, &data)
		if err != nil {
			log.Println("could not unmarshal json")
		}

		if data.User != user {
			log.Warningf("an invalidated user tried to make websocket request for %s, invalidated user: %s, user: %s", data.MsgType, data.User, user)
			continue
		}

		switch data.MsgType {
		// TODO: add support to remove message.
		case values.WebsocketOpenMsgType:
			handleLoadUserContent(user)

		case values.RequestMessages:
			msg.handleRequestMessages(user)

		case values.NewMessageMsgType:
			msg.handleNewMessage()

		case values.CreateRoomMsgType:
			msg.handleCreateNewRoom()

		case values.JoinRoomMsgType:
			msg.handleUserAcceptRoomRequest()

		case values.ExitRoomMsgType:
			msg.handleExitRoom(user)

		case values.RequestUsersToJoinRoomMsgType:
			msg.handleRequestUserToJoinRoom()

		case values.NewFileUploadMsgType:
			msg.handleNewFileUpload()

		case values.UploadFileChunkMsgType:
			msg.handleUploadFileChunk()

		case values.UploadFileSuccessMsgType:
			msg.handleUploadFileUploadComplete()

		case values.RequestDownloadMsgType:
			msg.handleRequestDownload(user)

		case values.DownloadFileChunkMsgType:
			msg.handleFileDownload(user)

		case values.StartClassSession:
			classSessions.startClassSession(msg, user)

		case values.JoinClassSession:
			classSessions.joinClassSession(msg, user)

		case values.EndClassSession:
			classSessions.endClassSession(user)

		case values.RenegotiateSDP:
			sdpConstruct{}.acceptRenegotiation(msg)

		case values.SearchUserMsgType:
			handleSearchUser(data.SearchText, user)

		default:
			log.Println("Could not convert required type", data.MsgType)
		}
	}
}
