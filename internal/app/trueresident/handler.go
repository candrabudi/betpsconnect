package trueresident

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

func (h *handler) GetTrueResidents(c *gin.Context) {
	ctx := c.Request.Context()
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil || offset < 0 {
		offset = 0
	}

	filter := dto.TrueResidentFilter{
		NamaKabupaten: c.Query("nama_kabupaten"),
		NamaKecamatan: c.Query("nama_kecamatan"),
		NamaKelurahan: c.Query("nama_kelurahan"),
		Jaringan:      c.Query("jaringan"),
		TPS:           c.Query("tps"),
		IsManual:      c.Query("is_manual"),
		Nama:          c.Query("nama"),
	}

	// Menggunakan nilai limit, offset, dan filter untuk memanggil service.GetListResident
	data, err := h.service.GetAll(ctx, limit, offset, filter, c.Value("user"))
	if err != nil {
		response := util.APIResponse("Failed to retrieve resident list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of residents", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Store(c *gin.Context) {
	var payload dto.TrueResidentPayload
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

	response := util.APIResponse("Success store valid resident", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	var payload dto.PayloadUpdateTrueResident
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

	err := h.service.Update(c, ID, payload)

	if err != nil {
		response := util.APIResponse(fmt.Sprintf("%s", err.Error()), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success update valid resident", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetTpsOnValidResident(c *gin.Context) {
	ctx := c.Request.Context()
	filter := dto.FindTpsByDistrict{
		NamaKabupaten: c.Query("nama_kabupaten"),
		NamaKecamatan: c.Query("nama_kecamatan"),
		NamaKelurahan: c.Query("nama_kelurahan"),
	}

	data, err := h.service.GetTpsOnValidResident(ctx, filter)
	if err != nil {
		response := util.APIResponse("Failed to retrieve tps : "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get data tps", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
