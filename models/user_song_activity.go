package models

type UserSongActivity struct {
	ID        int
	AccountID int
	Song      Song
	Latitude  float64
	Longitude float64
}
