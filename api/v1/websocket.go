package v1

import (
	"github.com/gofiber/fiber"
)

func ServeWs(c *fiber.Ctx) {
	//Run hub
	id := fn0(c)
	if id != nil {
		//  err := hub.Upgrader.Upgrade(c, func(conn *websocket.Conn) {
		// 	hub.HubConstruct.RegisterWS(conn, id)
		// }
		// if err != nil {
		// 	log.Errorln("error upgrading websocket, err:", err)
		// 	return
		// }
		return
	}
	// 	err := hub.Hub.upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
	// 		if err := conn.WriteJSON(struct {
	// 			MessageType string `json:"msgType"`
	// 		}{
	// 			// values.UnauthorizedAcces,
	// 		}); err != nil {
	// 			log.Errorln("could not send unauthorized message to user, err:", err)
	// 		}

	// 		if err := conn.Close(); err != nil {
	// 			log.Errorln("error closing websocket on unauthorized access, err:", err)
	// 		}
	// 	})
	//    if err != nil {
	// 		log.Errorln("error upgrading websocket while sending error message, err: ", err)
	// 		return
	// 	}

}
