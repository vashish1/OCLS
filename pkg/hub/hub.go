package hub

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fasthttp/websocket"

	"github.com/vashish1/OnlineClassPortal/pkg/models"
	"github.com/vashish1/OnlineClassPortal/vendor/github.com/metaclips/LetsTalk/backend/values"
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

func (h *models.Hub) RegisterWS(ws *websocket.Conn, email string) {
	c := &models.Connection{Send: make(chan []byte, 256), Ws: ws}

	s := models.Subscription{Conn: c, User: email}
	HubConstruct.Register <- s
	log.Infoln("user", email, "Connected")
	go s.readPump(email)
	go s.writePump()

}

func (h *models.Hub) Run() {
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
		case s := <-h.Register:
			h.Users.Mutex.Lock()

			if _, exists := h.Users.Users[s.User]; !exists {
				connections := make(map[*models.Connection]bool)

				h.Users.Users[s.User] = connections
			}

			h.Users.Users[s.User][s.Conn] = true

			log.Infoln(s.User, "registered")
			go broadcastOnlineStatusToAllUserRoom(s.User, true)
			h.Users.Mutex.Unlock()

		case s := <-h.UnRegister:
			h.Users.Mutex.Lock()
			connections, exists := h.Users.Users[s.User]
			if exists {
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
					close(s.Conn.Send)
					if len(connections) == 0 {
						delete(h.Users.Users, s.User)
						go broadcastOnlineStatusToAllUserRoom(s.User, false)
						log.Infoln(s.User, "offline")
					}
					log.Infoln(s.User, "subscription removed")
				}
			}
			h.Users.Mutex.Unlock()

		case m := <-h.Broadcast:
			h.Users.Mutex.RLock()
			connections := h.Users.Users[m.User]
			for c := range connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(connections, c)
					if len(connections) == 0 {
						h.Users.Mutex.Lock()
						delete(h.Users.Users, m.User)
						h.Users.Mutex.Unlock()
						go broadcastOnlineStatusToAllUserRoom(m.User, false)
					}
				}
			}
			h.Users.Mutex.RUnlock()
		}
	}
}

func (h *models.Hub) sendMessage(msg []byte, user string) {
	m := models.WsMessage{msg, user}
	HubConstruct.Broadcast <- m
}

// WritePump pumps messages from the hub to the websocket connection.
func (s *models.Subscription) writePump() {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
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
func (c *models.Connection) write(mt int, payload []byte) error {
	if err := c.Ws.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		return err
	}

	return c.Ws.WriteMessage(mt, payload)
}

// ReadPump pumps messages from the websocket connection to the hub.
func (s models.Subscription) readPump(user string) {
	c := s.Conn

	defer func() {
		HubConstruct.UnRegister <- s
		c.Ws.Close()
	}()

	c.Ws.SetReadDeadline(time.Now().Add(pongWait))

	c.Ws.SetPongHandler(
		func(string) error {
			return c.Ws.SetReadDeadline(time.Now().Add(pongWait))
		})

	for {
		var err error
		var msg messageBytes
		_, msg, err = c.Ws.ReadMessage()

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

		// 		switch data.MsgType {
		// 		// TODO: add support to remove message.
		// 		case values.WebsocketOpenMsgType:
		// 			handleLoadUserContent(user)

		// 		case values.RequestMessages:
		// 			msg.handleRequestMessages(user)

		// 		case values.NewMessageMsgType:
		// 			msg.handleNewMessage()

		// 		case values.CreateRoomMsgType:
		// 			msg.handleCreateNewRoom()

		// 		case values.JoinRoomMsgType:
		// 			msg.handleUserAcceptRoomRequest()

		// 		case values.ExitRoomMsgType:
		// 			msg.handleExitRoom(user)

		// 		case values.RequestUsersToJoinRoomMsgType:
		// 			msg.handleRequestUserToJoinRoom()

		// 		case values.NewFileUploadMsgType:
		// 			msg.handleNewFileUpload()

		// 		case values.UploadFileChunkMsgType:
		// 			msg.handleUploadFileChunk()

		// 		case values.UploadFileSuccessMsgType:
		// 			msg.handleUploadFileUploadComplete()

		// 		case values.RequestDownloadMsgType:
		// 			msg.handleRequestDownload(user)

		// 		case values.DownloadFileChunkMsgType:
		// 			msg.handleFileDownload(user)

		// 		case values.StartClassSession:
		// 			classSessions.startClassSession(msg, user)

		// 		case values.JoinClassSession:
		// 			classSessions.joinClassSession(msg, user)

		// 		case values.EndClassSession:
		// 			classSessions.endClassSession(user)

		// 		case values.RenegotiateSDP:
		// 			sdpConstruct{}.acceptRenegotiation(msg)

		// 		case values.SearchUserMsgType:
		// 			handleSearchUser(data.SearchText, user)

		// 		default:
		// 			log.Println("Could not convert required type", data.MsgType)
		// 		}
		// 	}
		// }
	}
}
