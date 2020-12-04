package v1

import (
	"github.com/gofiber/fiber"
	"github.com/vashish1/OnlineClassPortal/pkg/hub"
)

func ServeWs(c *fiber.Ctx) {
	//Run hub
	id := fn0(c)
	if id != nil {
		err := hub.Upgrader.Upgrade(c, func(conn *websocket.Conn) {
			hub.HubConstruct.RegisterWS(conn, id)
		})
		if err != nil {
			log.Errorln("error upgrading websocket, err:", err)
			return
		}
		return
	}

	////Is this even required?here or should i sinply return error after first upgrade, and what is the need of upgrading it now.
		err := hub.upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
			if err := conn.WriteJSON(struct {
				MessageType string `json:"msgType"`
			}{
				models.UnauthorizedAccess,
			}); err != nil {
				log.Errorln("could not send unauthorized message to user, err:", err)
			}

			if err := conn.Close(); err != nil {
				log.Errorln("error closing websocket on unauthorized access, err:", err)
			}
		})
	   if err != nil {
			log.Errorln("error upgrading websocket while sending error message, err: ", err)
			return
		}

}
