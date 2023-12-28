package resident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/pkg/util"
	"fmt"
	"io"
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

func (h *handler) GetResidents(c *gin.Context) {
	ctx := c.Request.Context()

	// Mendapatkan nilai limit dan offset dari permintaan HTTP
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil || limit <= 0 {
		limit = 20 // Nilai default jika limit tidak valid atau tidak ada dalam permintaan
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil || offset < 0 {
		offset = 0 // Nilai default jika offset tidak valid atau tidak ada dalam permintaan
	}

	// Membuat DTO ResidentFilter dari nilai yang diterima dari permintaan HTTP
	filter := dto.ResidentFilter{
		NamaKabupaten: c.Query("nama_kabupaten"),
		NamaKecamatan: c.Query("nama_kecamatan"),
		NamaKelurahan: c.Query("nama_kelurahan"),
		TPS:           c.Query("tps"),
		Nama:          c.Query("nama"),
	}

	// Menggunakan nilai limit, offset, dan filter untuk memanggil service.GetListResident
	data, err := h.service.GetListResident(ctx, limit, offset, filter)
	if err != nil {
		response := util.APIResponse("Failed to retrieve resident list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of residents", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetGroupBy(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.service.GetListResidentGroup(ctx)
	if err != nil {
		response := util.APIResponse("Failed to retrieve resident list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of residents", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Store(c *gin.Context) {
	var payload dto.PayloadStoreResident
	if err := c.ShouldBind(&payload); err != nil {
		errorMessage := gin.H{"errors": "Please fill data"}
		if err != io.EOF {
			errors := util.FormatValidationError(err)
			errorMessage = gin.H{"errors": errors}
		}
		response := util.APIResponse("Error validation", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.service.Store(c, payload)

	if err != nil {
		response := util.APIResponse(fmt.Sprintf("%s", err.Error()), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success store server", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
