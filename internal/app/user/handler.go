package user

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/pkg/constants"
	"betpsconnect/pkg/util"
	"io"
	"net/http"
	"regexp"

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

func (h *handler) LoginUser(c *gin.Context) {
	var payload dto.PayloadLogin
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

	resultLogin, err := h.service.LoginUser(c, payload)

	if err == constants.UserNotFound {
		response := util.APIResponse(err.Error(), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success login user", http.StatusOK, "success", resultLogin)
	c.JSON(http.StatusOK, response)
}

func (h *handler) LogoutUser(c *gin.Context) {
	header := c.Request.Header["Authorization"]
	rep := regexp.MustCompile(`(Bearer)\s?`)
	bearerStr := rep.ReplaceAllString(header[0], "")
	err := h.service.Logout(c, bearerStr)

	if err == constants.UserNotFound {
		response := util.APIResponse(err.Error(), http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := util.APIResponse("Success login user", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
