package websocket


import (
	"time"
	"github.com/gorilla/websocket"
	"log"
	"github.com/prongbang/next/structs"
	"github.com/prongbang/next/dao"
	"github.com/prongbang/next/utils"
	"encoding/json"
	"fmt"
)

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil
	})
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		h.broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(dao *dao.DaoConfig, mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))

	var response []byte

	// save event to Mgo
	if mt == websocket.TextMessage {

		// ADD
		var event structs.Event
		json.Unmarshal(payload, &event)
		fmt.Println(event)
		if event.Message != "" {
			dao.Save(event, dao.COLLECTION_EVENT)
			var res = structs.Response{101}
			response = utils.Type2JsonByte(res)
		}

		// Find
		var req structs.Request
		json.Unmarshal(payload, &req)
		if req != (structs.Request{}) {
			//response = Type2JsonByte(FindByLatLng(req))
			response = utils.Type2JsonByte(dao.FindById(1))
		}
	}

	return c.ws.WriteMessage(mt, response)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump(dao *dao.DaoConfig) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(dao, websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(dao, websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(dao, websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
