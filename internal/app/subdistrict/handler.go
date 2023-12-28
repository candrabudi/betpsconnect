package subdistrict

import (
	"betpsconnect/internal/dto"
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

func (h *handler) GetByDistrict(c *gin.Context) {
	ctx := c.Request.Context()
	filter := dto.GetByDistrict{
		NamaKecamatan: c.Query("nama_kecamatan"),
		NamaKelurahan: c.Query("nama_kelurahan"),
	}

	data, err := h.service.GetByDistrict(ctx, filter)
	if err != nil {
		response := util.APIResponse("Failed to retrieve district list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of districts", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
