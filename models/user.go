package models

import "time"

type User struct {
	AccountID int64
	Name      string
	BirthDay  time.Time
	Photos    []UserPhoto
}
