package resident

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.GET("/list", h.GetResidents)
	g.GET("/list/groupby", h.GetGroupBy)
	g.POST("/store", h.Store)
}
