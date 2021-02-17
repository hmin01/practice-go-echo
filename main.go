package main

import (
	"strconv"
	app "practice-go-echo/routes"
)

const (
	PORT = 4000
)

func main() {
	router := app.Router()
	router.Logger.Fatal(router.Start(":" + strconv.Itoa(PORT)))
}