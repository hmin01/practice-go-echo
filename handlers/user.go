package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"
	// Echo
	echo "github.com/labstack/echo"
	// Modules (custom)
	model "practice-go-echo/models"
	util "practice-go-echo/utils"
)

func UserSignup(ctx echo.Context) error {
	inName, err := url.QueryUnescape(ctx.FormValue("name"))
	if eval := util.CatchError(err); eval != nil {
		return eval
	}
	inEmail, err := url.QueryUnescape(ctx.FormValue("email"))
	if eval := util.CatchError(err); eval != nil {
		return eval
	}
	inPassword, err := url.QueryUnescape(ctx.FormValue("password"))
	if eval := util.CatchError(err); eval != nil {
		return eval
	}

	log.Print(inName, inEmail, inPassword)
	var strCount string
	// Duplicate check for name
	err = db.QueryRow("SELECT count(*) FROM user WHERE user_name=?", inName).Scan(&strCount)
	if eval := util.CatchError(err); eval != nil {
		return eval
	}
	count, err := strconv.Atoi(strCount)
	if eval := util.CatchError(err); eval != nil {
		return eval
	} else if count > 0 {
		return util.HTTPResponse(ctx, false, 12, []interface{}{"Username already exists."})
	}
	// Duplicate check for email
	err = db.QueryRow("SELECT count(*) FROM user WHERE user_email=?", inEmail).Scan(&strCount)
	if eval := util.CatchError(err); eval != nil {
		return eval
	}
	count, err = strconv.Atoi(strCount)
	if eval := util.CatchError(err); eval != nil {
		return eval
	} else if count > 0 {
		return util.HTTPResponse(ctx, false, 12, []interface{}{"This email already exists."})
	}

	// Insert user (sign up)
	result, err := db.Exec("INSERT INTO user (user_name, user_email, user_password) VALUE (?, ?, ?)", inName, inEmail, inPassword)
	if eval := util.CatchError(err); eval != nil {
		return eval
	}
	affected, err := result.RowsAffected()
	if eval := util.CatchError(err); eval != nil {
		return eval
	} else if affected == 0 {
		return util.HTTPResponse(ctx, false, 10, []interface{}{"No columns were affected."})
	} else {
		return util.HTTPResponse(ctx, true, 0, []interface{}{"Welcome to the membership. Please sign-in for use."})
	}
}

func UserSignin(ctx echo.Context) error {
	inEmail, err := url.QueryUnescape(ctx.FormValue("email"))
	if eval := util.CatchError(err); eval != nil {
		return eval
	}
	inPassword, err := url.QueryUnescape(ctx.FormValue("password"))
	if eval := util.CatchError(err); eval != nil {
		return eval
	}

	// Create user info structure
	var strTime string
	var info model.UserInfo
	// Select user
	err = db.QueryRow("SELECT user_uuid, user_email, user_name, reg_date FROM user WHERE user_email=? AND user_password=?", inEmail, inPassword).Scan(&info.Uuid, &info.Email, &info.Name, &strTime)
	if err == sql.ErrNoRows {
		return util.HTTPResponse(ctx, false, 11, []interface{}{"This email is not subscribed or the password in invalid."})
	} else if eval := util.CatchError(err); eval != nil {
		return eval
	} else {
		// Convert string to time
		info.RegDate, err = util.ConvertToTime(strTime)
		if eval := util.CatchError(err); eval != nil {
			return eval
		}
		fmt.Println(info)
		// Return
		return util.HTTPResponse(ctx, true, 0, nil)
	}
}