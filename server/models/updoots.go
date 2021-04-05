package models

type Updoot struct {
	UserID int `gorm:"primaryKey"`
	PostID int `gorm:"primaryKey"`
	Value  int
}
