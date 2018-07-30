package models

type Song struct {
	ID      int
	AlbumID Album
	Album   Album
	Artists []Artist `gorm:"many2many:song_artists"`
	Name    string
}
