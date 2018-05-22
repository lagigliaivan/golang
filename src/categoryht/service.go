package categoryht

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.GET("/ping", Get)
	return router
}

func Get(c *gin.Context) {
	c.String(200, "pong")
}
