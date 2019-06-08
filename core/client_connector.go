package core

import (
	"HomeServices/nlp"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/http"
)

var upgrader = websocket.Upgrader{ EnableCompression: true, CheckOrigin: func(r *http.Request) bool {
	return true
}}

// ClientConnector - connects with browser client
type ClientConnector struct {
	connection *websocket.Conn
	RemoteAddress net.Addr
	processor *nlp.Processor
	session *ClientSession
	server *Server
}

// NewClientConnector - create and return new client WS connection
func NewClientConnector(w http.ResponseWriter, r *http.Request, p *nlp.Processor, s *Server) (*ClientConnector, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade to WS protocol: %v", err)
	}

	c := &ClientConnector{ws, ws.RemoteAddr(), p, NewClientSession(p), s}

	go c.readRequest()
	return c, nil
}

// readRequest - loop for receiving client input
func (c *ClientConnector) readRequest() {
	defer c.connection.Close()

	var message Message
	for {
		err := c.connection.ReadJSON(&message)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("ClientConnector(%v): failed to read JSON request: %v", c.RemoteAddress, err)
			break
		}

		text := fmt.Sprintf("%v", message.Text)

		answer := c.session.handleRequest(Message{text})
		err = c.write(answer)
		if err != nil {
			log.Printf("ClientConnector(%v): failed to handle write response: %v", c.RemoteAddress, err)
		}

		if c.session.IsCompleted {
			c.session = NewClientSession(c.processor)
		}
	}

	log.Printf("ClientConnector(%v): client disconnected", c.RemoteAddress)
	c.server.OnClientDisconnected(c)
}

func findMostProbably(entities []nlp.Entity) (entity nlp.Entity) {
	for _, e := range entities {
		if e.Confidence > entity.Confidence {
			entity = e
		}
	}

	return
}

func (c *ClientConnector) write(answer Answer) error {
	err := c.connection.WriteJSON(answer)
	return err
}

