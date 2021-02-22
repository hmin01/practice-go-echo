package utils

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	echo "github.com/labstack/echo"
	// Modules
	response "practice-go-echo/models"
)

func CatchError(err error) error {
	if err != nil {
		PrintLog("error", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		return nil
	}
}

func HTTPResponse(ctx echo.Context, result bool, code int, message []interface{}) error {
	format := response.ResponseFormat{
		Result: result,
		Code: code,
		Message: message,
	}
	return ctx.JSON(http.StatusOK, format)
}

func PrintLog(logType string, message string) {
	var buffer bytes.Buffer
	switch logType {
		case "error":
			buffer.WriteString("\u001B[31m[ERROR] ")
		case "warning":
			buffer.WriteString("\u001B[33m[WARNING] ")
		case "notice":
			buffer.WriteString("\u001B[36m[NOTICE] ")
		case "debug":
		default:
			buffer.WriteString("\u001B[35m[ERROR] ")
	}
	buffer.WriteString(message)
	buffer.WriteString("\u001B[0m")
	log.Print(buffer.String())
}

func OccurFatalError(message string) {
	var buffer bytes.Buffer
	buffer.WriteString("\u001B[31m[ERROR] ")
	buffer.WriteString(message)
	buffer.WriteString("\u001B[0m")
	log.Fatal(buffer.String())
}

func CreateError(message string) error {
	return errors.New(message)
}