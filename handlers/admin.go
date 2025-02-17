package handlers

import (
	"delivery/entities"
	httppkg "delivery/pkg/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateAPI(c *gin.Context) {
	var body entities.API
	err := c.ShouldBindJSON(&body)
	if err != nil {
		bodyErr := handleBodyParseError(err)
		h.handleResponse(c, httppkg.BadRequest,
			fmt.Sprintf("field: %v  message: %v", bodyErr.Field, bodyErr.Message), nil)
		return
	}
	body.ID = uuid.NewString()
	err = h.adminController.CreateAPI(c, &body)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error(), nil)
		return
	}

	h.handleResponse(c, httppkg.OK, "", body)
}

func (h *Handler) CreateAPIDetails(c *gin.Context) {
	var body entities.APIDetails
	err := c.ShouldBindJSON(&body)
	if err != nil {
		bodyErr := handleBodyParseError(err)
		h.handleResponse(c, httppkg.BadRequest,
			fmt.Sprintf("field: %v  message: %v", bodyErr.Field, bodyErr.Message), nil)
		return
	}
	body.ID = uuid.NewString()
	err = h.adminController.CreateAPIDetails(c, &body)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error(), nil)
		return
	}

	h.handleResponse(c, httppkg.OK, "", body)
}

func (h *Handler) ExecuteAPI(c *gin.Context) {
	apiName := c.Param("api_name")
	method := c.Request.Method

	fmt.Println(apiName, method)
	

	res, err := h.adminController.ExecuteAPI(c, apiName, method)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error(), nil)
		return
	}

	h.handleResponse(c, httppkg.OK, "", res)
}
