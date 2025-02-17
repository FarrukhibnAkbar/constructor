package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"delivery/configs"
	adminController "delivery/controllers/admin"
	"delivery/logger"
	e "delivery/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	httppkg "delivery/pkg/http"
)

type Handler struct {
	cfg             *configs.Configuration
	log             logger.LoggerI
	adminController adminController.AdminController
	redis           *redis.Client
}

func New(
	cfg *configs.Configuration,
	log logger.LoggerI,
	adminController adminController.AdminController,
	redis *redis.Client,
) Handler {
	return Handler{
		cfg:             cfg,
		log:             log,
		adminController: adminController,
		redis:           redis,
	}
}

// Response - response JSON formatida qaytaruvchi struct
type Response struct {
	Timestamp string      `json:"timestamp"`
	Status    int         `json:"status"`
	Message   *string     `json:"message"`
	Error     *string     `json:"error"`
	Path      string      `json:"path"`
	Data      interface{} `json:"data"`
}

// handleResponse - o'zgarmas response qaytarish uchun tayyorlangan funksiya,
// Ehtiyot bo'lish kerak: bu universal funksiya bo'lib errorni ham success responseni ham qaytariladi
// Error yuborilganda errMessagega error message yoziladi dataga esa nil yuboriladi.
// Success message yuborilganda errMessage ga bo'sh string(yani "") yuboriladi.
func (h *Handler) handleResponse(c *gin.Context, status httppkg.Status, errMessage string, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}
	var error *string
	error = &status.Status
	if status.Code >= 200 && status.Code < 300 {
		error = nil
	}

	c.JSON(status.Code, Response{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    status.Code,
		Message:   &errMessage,
		Error:     error,
		Data:      data,
		Path:      c.Request.URL.Path,
	})
}

type ResponseForExec struct {
	Info string `json:"info"`
	ID   string `json:"id"`
}

// StatusFromError ...
func StatusFromError(err error) httppkg.Status {
	if err == nil {
		return httppkg.OK
	}

	code, ok := e.ExtractStatusCode(err)
	if !ok || code == http.StatusInternalServerError {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	} else if code == http.StatusNotFound {
		return httppkg.Status{
			Code:        http.StatusNotFound,
			Status:      "NOT_FOUND",
			Description: err.Error(),
		}
	} else if code == http.StatusBadRequest {
		return httppkg.Status{
			Code:        http.StatusBadRequest,
			Status:      "BAD_REQUEST",
			Description: err.Error(),
		}
	} else if code == http.StatusForbidden {
		return httppkg.Status{
			Code:        http.StatusForbidden,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else if code == http.StatusUnauthorized {
		return httppkg.Status{
			Code:        http.StatusUnauthorized,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	}

}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func handleBodyParseError(err error) ValidationError {
	var validationErrors ValidationError

	var ute *json.UnmarshalTypeError
	if errors.As(err, &ute) {
		// line, column := findLineAndColumn(body, ute.Offset)
		validationErrors = ValidationError{
			Field:   ute.Field,
			Message: fmt.Sprintf("Type mismatch: expected %s but got %s", ute.Type.String(), ute.Value),
		}
	} else if syntaxErr, ok := err.(*json.SyntaxError); ok {
		validationErrors = ValidationError{
			Field:   "JSON",
			Message: fmt.Sprintf("Syntax error at offset %d", syntaxErr.Offset),
		}
	} else {
		validationErrors = ValidationError{
			Field:   "Unknown",
			Message: err.Error(),
		}
	}

	return validationErrors
}
