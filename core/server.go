package core

import (
	"HomeServices/nlp"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	clients []*ClientConnector
	processor *nlp.Processor
}

func NewServer() (*Server, error) {
	p := nlp.NewProcessor()
	return &Server{processor: p}, nil
}

func (s *Server) Listen(address string, port int) error {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		client, err := NewClientConnector(w, r, s.processor, s)
		if err != nil {
			log.Printf("%v request failed: failed to create client connector: %v", "/send", err)
			return
		}

		s.clients = append(s.clients, client)

		log.Printf("New client connected(%v)", client.RemoteAddress)
	})

	err := http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), nil)
	return err
}

func (s *Server) OnClientDisconnected(c *ClientConnector) {
	var id int
	for n, client := range s.clients {
		if client == c {
			id = n
			break
		}
	}

	s.clients = append(s.clients[:id], s.clients[id+1:]...)
}


