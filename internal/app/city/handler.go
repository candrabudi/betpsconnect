package city

import (
	"betpsconnect/internal/factory"
	"betpsconnect/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) GetCity(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := h.service.GetCity(ctx)
	if err != nil {
		response := util.APIResponse("Failed to retrieve city list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of cities", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
