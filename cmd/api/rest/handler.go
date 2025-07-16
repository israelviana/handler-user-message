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
	r.router.GET("/spiral", r.Spiral)
}

func (r *RestHandler) Spiral(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "O símbolo da espiral é comumente considerado o sinal da vida e representa o ciclo da vida como um todo, desde o nascimento. vida, morte e, finalmente, renascimento. Use a espiral em seu ofício para reverenciar e honrar a vida, pois as espirais podem ser encontradas em muitos lugares da natureza (galáxias e conchas).",
	})
}
