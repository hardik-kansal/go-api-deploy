package api

import (
	"github.com/gin-gonic/gin"
	_"github.com/gin-gonic/gin/binding"
	db "github.com/hardik-kansal/go-api-deploy/db/sqlc"
	"github.com/hardik-kansal/go-api-deploy/util"
)

type Server struct {
	config     util.ConfigPwd
	store      *db.SQLStore
	router     *gin.Engine
}

func NewServer(config util.ConfigPwd, store *db.SQLStore) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}