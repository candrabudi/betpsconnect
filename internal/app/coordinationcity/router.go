package coordinationcity

import (
	"betpsconnect/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.Use(middleware.Authenticate())
	g.GET("/list", h.GetListCoordinationCity)
	g.POST("/store", h.Store)
	g.PUT("/update/:id", h.Update)
}
