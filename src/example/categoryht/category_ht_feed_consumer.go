package categoryht

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type Order struct {
	Id int `json:"order_id"`
}

type Handler struct {
}

func (*Handler) Get(c *gin.Context) {
	c.String(200, "pong")
}

func (*Handler) Post(c *gin.Context) {

	if c.Request.Body == nil {
		c.String(400, "order_id cannot be null")
	} else {

		order := Order{}
		c.BindJSON(&order)
		fmt.Printf("Receiving POST for orderId: %d\n", order.Id)
	}
}

func SetupRouter() *gin.Engine {

	handler := Handler{}
	router := gin.Default()

	router.GET("/ping", handler.Get)
	router.POST("/category/handlingtime", handler.Post)

	return router
}

func setHandler(handler *Handler) *gin.Engine {

	router := gin.Default()

	router.GET("/ping", handler.Get)
	router.POST("/category/handlingtime", handler.Post)

	return router
}
