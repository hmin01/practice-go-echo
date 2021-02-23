package utils

import (
	"encoding/json"
	// Echo
	"github.com/labstack/echo"
	// Session
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
)

func CreateSession(ctx echo.Context, sessionName string, expires int) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}
	// Set options
	sess.Options = &sessions.Options{
		Path: "/",
		MaxAge: expires,
		HttpOnly: true,
	}
	return nil
}

func SaveSession(ctx echo.Context, sessionName string, key string, data interface{}, serialize bool) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}
	// json marshal을 이용한 변환 여부
	if serialize {
		buf, err := json.Marshal(data)
		if err != nil {
			return err
		}
		sess.Values[key] = buf
	} else {
		sess.Values[key] = data
	}
	sess.Save(ctx.Request(), ctx.Response())
	return nil
}

func ReadSession(ctx echo.Context, sessionName string, key string, serialize bool) (map[string]interface{}, error) {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return nil, err
	}
	// json marshal을 이용한 변환 여부
	var info map[string]interface{}
	if serialize {
		err := json.Unmarshal([]byte(sess.Values[key].([]byte)), &info)
		if err != nil {
			return nil, err
		}
	} else {
		info[key] = sess.Values[key]
	}
	return info, nil
}