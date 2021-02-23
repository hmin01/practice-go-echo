package handlers

import (
	"log"
	"net/http"
	// Echo
	"github.com/labstack/echo"
	// Session
	util "practice-go-echo/utils"
)

const (
	SESS_NAME = "test#test"
)

func CreateSession(ctx echo.Context) error {
	// Get session
	err := util.CreateSession(ctx, SESS_NAME, 3600)
	if err != nil {
		return err
	}
	// Save data in session
	err = util.SaveSession(ctx, SESS_NAME, "user", "haemin", false)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, "Save session")
}

func ReadSession(ctx echo.Context) error {
	// Get session
	data, err := util.ReadSession(ctx, SESS_NAME, "user", false)
	if err != nil {
		return err
	}
	// Get data
	log.Print("Get session")
	log.Print(data)
	return ctx.String(http.StatusOK, "Read session")
}