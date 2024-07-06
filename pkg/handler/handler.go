package handler

import (
	"net/http"

	service "github.com/end1essrage/dndhelper-spellbook/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello!")
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/", h.sayHello)
		api.GET("/spell-book", h.GetSpellInfo)
	}

	return router
}
