package handlers

import (
	"log"
	"net/http"
	"net/url"
	// Echo
	echo "github.com/labstack/echo"
	// Model
	model "practice-go-echo/models"
)

func UserSignup(ctx echo.Context) error {
	inName, err := url.QueryUnescape(ctx.FormValue("name"))
	if err != nil {
		return err
	}
	inEmail, err := url.QueryUnescape(ctx.FormValue("email"))
	if err != nil {
		return err
	}
	inPassword, err := url.QueryUnescape(ctx.FormValue("password"))
	if err != nil {
		return err
	}

	log.Print(inName, inEmail, inPassword)

	// Response
	message := model.ResponseFormat{
		Result: true,
		Code: 0,
		Message: []interface{}{"OK"},
	}
	return ctx.JSON(http.StatusOK, message)
}