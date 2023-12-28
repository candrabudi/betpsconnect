package district

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.GET("/by-city", h.GetDistrictByCity)
}
