package api

import (
	"log"

	db "github.com/didate/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Serves HTTP Requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.PUT("/accounts", server.updateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	log.Println(err)
	return gin.H{"error": err.Error()}
}
