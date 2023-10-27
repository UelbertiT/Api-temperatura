package server

import (
	"api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type server struct {
	port   string
	server *gin.Engine
}

func NewSever() server {
	return server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}
