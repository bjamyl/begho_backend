package api

import (
	"log"

	db "github.com/bjamyl/begho/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}

	server.router.POST("/users", server.createUser)
	server.router.GET("/users", server.fetchUsers)

	return server
}

func (server *Server) Run(address string) {
	err := server.router.Run(address)
	if err != nil {
		log.Fatal("could not start this server: ", err.Error())
		return
	}
}
