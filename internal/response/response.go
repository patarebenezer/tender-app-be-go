package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorData struct {
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Envelope struct {
	Status      string            `json:"status"`
	Code        int               `json:"code"`
	Message     string            `json:"message"`
	Total       int               `json:"total"`
	Data        interface{}       `json:"data"`
	ErrorFields map[string]string `json:"error_fields"`
	ErrorData   ErrorData         `json:"error_data"`
}

func OK(c *gin.Context, data interface{}, total int, message string) {
	c.JSON(http.StatusOK, Envelope{
		Status:      "Success",
		Code:        200,
		Message:     message,
		Total:       total,
		Data:        data,
		ErrorFields: map[string]string{},
		ErrorData:   ErrorData{},
	})
}

func Fail(c *gin.Context, httpCode int, message string, fields map[string]string) {
	c.JSON(httpCode, Envelope{
		Status:      "Fail",
		Code:        400,
		Message:     message,
		Total:       0,
		Data:        nil,
		ErrorFields: fields,
		ErrorData:   ErrorData{Name: "Error", Message: message, Data: nil},
	})
}
