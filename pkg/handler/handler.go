package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello!")
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	api.GET("/", h.sayHello)

	return router
}
