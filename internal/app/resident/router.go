package resident

import (
	"betpsconnect/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.Use(middleware.Authenticate())
	g.GET("/list", h.GetResidents)
	g.GET("/list/groupby", h.GetGroupBy)
	g.GET("/detail/:resident_id", h.DetailResident)
	g.GET("/tps/subdistrict", h.GetTpsBySubDistrict)
	g.POST("/store", h.Store)
	g.POST("/validate", h.ValidateResident)
}
