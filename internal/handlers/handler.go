package handlers

import (
	"github.com/gin-gonic/gin"
	"shina/internal/site"
	"shina/internal/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		utils.RenderHTML(c.Writer, c.Request, "home", &site.HTMLData{})
	})

	return router
}
