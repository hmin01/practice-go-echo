package handlers

import (
	"net/http"
	// Echo
	echo "github.com/labstack/echo"
	// Model
	model "practice-go-echo/models"
)

func init() {
	// Create db pool
}

func HealthCheck(ctx echo.Context) error {
	response := model.ResponseFormat{
		Result: true,
		Code: 100, 
		Message: []interface{}{"alive"},
	}
	return ctx.JSON(http.StatusOK, response)
}