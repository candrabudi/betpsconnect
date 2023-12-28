package subdistrict

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.GET("/by-district", h.GetByDistrict)
}
