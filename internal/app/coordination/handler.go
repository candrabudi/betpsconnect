package coordinationcity

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/pkg/util"
	"net/http"
	"strconv"

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

func (h *handler) GetListCoordinationCity(c *gin.Context) {
	ctx := c.Request.Context()
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil || offset < 0 {
		offset = 0
	}

	filter := dto.ResidentFilter{
		Nama: c.Query("nama"),
	}

	data, err := h.service.GetListCoordinationCity(ctx, limit, offset, filter, c.Value("user"))
	if err != nil {
		response := util.APIResponse("Failed to retrieve coordination city list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of coordination city", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
