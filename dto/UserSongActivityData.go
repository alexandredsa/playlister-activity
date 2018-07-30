package dto

type UserSongActivityData struct {
	UserID    int     `json:"user_id"`
	Artist    string  `json:"artist"`
	Album     string  `json:"album"`
	Track     string  `json:"track"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
