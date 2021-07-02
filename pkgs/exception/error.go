package exception

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type FailResponse struct {
	Status string            `json:"status"`
	Data   map[string]string `json:"data"`
}

// Handle 500 error and 404 fail
func NotFoundError(err error, c *gin.Context) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, FailResponse{
			Status: "fail",
			Data:   nil,
		})
	} else {
		c.JSON(500, ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
}

// Handle 500 error
func CommonError(err error, c *gin.Context) {
	c.JSON(500, ErrorResponse{
		Status:  "error",
		Message: err.Error(),
	})
}
