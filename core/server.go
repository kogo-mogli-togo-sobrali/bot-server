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
		client, err := NewClientConnector(w, r, s.processor)
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


