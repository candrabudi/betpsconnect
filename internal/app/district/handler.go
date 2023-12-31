package district

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

func (h *handler) GetDistrictByCity(c *gin.Context) {
	ctx := c.Request.Context()
	filter := dto.GetByCity{
		NamaKabupaten: c.Query("nama_kabupaten"),
		NamaKecamatan: c.Query("nama_kecamatan"),
	}

	data, err := h.service.GetDistrictByCity(ctx, filter, c.Value("user"))
	if err != nil {
		response := util.APIResponse("Failed to retrieve district list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of districts", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
