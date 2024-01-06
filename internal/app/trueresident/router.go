package trueresident

import (
	"betpsconnect/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.Use(middleware.Authenticate())
	g.GET("/list", h.GetTrueResidents)
	g.POST("/store", h.Store)
	g.PUT("/update/:id", h.Update)
	g.GET("/tps/subdistrict", h.GetTpsOnValidResident)
}
