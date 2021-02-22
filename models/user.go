package models

import (
	"time"
)

type UserInfo struct {
	Uuid 		int				`json:"uuid" xml:"uuid"`
	Name 		string		`json:"name" xml:"name"`
	Email 	string		`json:"email" xml:"email"`
	RegDate time.Time	`json:"regDate" xml:"regDate"`
}