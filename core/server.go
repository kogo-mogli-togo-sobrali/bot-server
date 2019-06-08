package core

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {

}

func NewServer() (*Server, error) {
	s := &Server{}

	err := s.Listen(viper.GetString("server.address"), viper.GetInt("server.port"))
	if err != nil {
		return nil, fmt.Errorf("failed to start server: %v", err)
	}


}

func (s *Server) Listen(address string, port int) error {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {

	})

	return nil, &ClientConnector{}
}


