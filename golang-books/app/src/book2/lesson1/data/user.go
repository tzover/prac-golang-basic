package data

import "time"

type Session struct {
	Id      int
	Uuid    string
	Email   string
	UserId  int
	Created time.Time
}
