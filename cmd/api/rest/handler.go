package rest

import (
	"github.com/gin-gonic/gin"
)

type IRestHandler interface {
	InitRestHandler()
	Spiral(c *gin.Context)
}

type RestHandler struct {
	router *gin.Engine
}

func NewRestHandler(router *gin.Engine) *RestHandler {
	return &RestHandler{
		router: router,
	}
}

func (r *RestHandler) InitRestHandler() {
	r.router.GET("/health", r.health)
}

func (r *RestHandler) health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "play.sospita-craft.me",
	})
}
