package coordinationtps

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/pkg/util"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

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

func (h *handler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil || limit <= 0 {
		limit = 20
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil || offset < 0 {
		offset = 0
	}

	filter := dto.CoordinationTpsFilter{
		Nama:          c.Query("nama"),
		NamaKabupaten: c.Query("nama_kabupaten"),
		NamaKecamatan: c.Query("nama_kecamatan"),
		NamaKelurahan: c.Query("nama_kelurahan"),
		Jaringan:      c.Query("jaringan"),
		Tps:           c.Query("tps"),
	}

	data, err := h.service.GetAll(ctx, limit, offset, filter, c.Value("user"))
	if err != nil {
		response := util.APIResponse("Failed to retrieve coordination tps list: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := util.APIResponse("Success get list of coordination tps", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Store(c *gin.Context) {
	var payload dto.PayloadStoreCoordinatorTps
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

	response := util.APIResponse("Success store coordination tps", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	var payload dto.PayloadUpdateCoordinatorTps
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

	response := util.APIResponse("Success update coordination tps", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	if ID == 0 {
		response := util.APIResponse("Please input ID coordination TPS ", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := h.service.Delete(c, ID)

	if err != nil {
		response := util.APIResponse(fmt.Sprintf("%s", err.Error()), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success delete coordination tps", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *handler) Export(c *gin.Context) {
	ctx := c.Request.Context()
	filter := dto.CoordinationTpsFilter{
		Nama:          c.Query("nama"),
		NamaKabupaten: c.Query("nama_kabupaten"),
		NamaKecamatan: c.Query("nama_kecamatan"),
		NamaKelurahan: c.Query("nama_kelurahan"),
		Tps:           c.Query("tps"),
		Jaringan:      c.Query("jaringan"),
	}
	data, err := h.service.Export(ctx, filter, c.Value("user"))
	if err != nil {
		response := util.APIResponse("Failed to retrieve export coordination tps: "+err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	fileName := fmt.Sprintf("export_kortps_%s.xlsx", time.Now().Format("20060102150405"))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}
