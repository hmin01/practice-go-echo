package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"
	// Echo
	echo "github.com/labstack/echo"
	// Project modules
	model "practice-go-echo/models"
	util "practice-go-echo/utils"
	// Database driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_TYPE = "mysql"
	DB_CONFIG = "./resources/init/db_config.json"
)

var db *sql.DB

func init() {
	// Get database config info
	data, result := util.ReadFile(DB_CONFIG)
	if !result {
		util.OccurFatalError("database config invalid.")
	}
	// Convert to json
	config := util.ConvertToJSON(data)
	// Create DSN
	dsn := config["user"].(string) + ":" + config["password"].(string) + "@tcp(" + config["host"].(string) + ":" + strconv.FormatFloat(config["port"].(float64), 'f', -1, 64) + ")/" + config["database"].(string)
	// Create database object
	var err error
	db, err = sql.Open(DB_TYPE, dsn)
	if err != nil {
		util.OccurFatalError(err.Error())
	}
	// Test connection
	connStat := db.Ping()
	if connStat != nil {
		util.OccurFatalError(connStat.Error())
	} else {
		// Set connection options
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(16)
		db.SetMaxIdleConns(16)
		util.PrintLog("notice", "Database connection successful.");
	}
}

func HealthCheck(ctx echo.Context) error {
	response := model.ResponseFormat{
		Result: true,
		Code: 100, 
		Message: []interface{}{"alive"},
	}
	return ctx.JSON(http.StatusOK, response)
}