package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phlucasfr/api_go/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("server is runnig at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
