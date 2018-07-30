package models

type Album struct {
	ID      int
	Artists []Artist `gorm:"many2many:song_artists"`
	Name    string
}
