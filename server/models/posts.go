package models

import (
	"time"
)

// post model
type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	Points    int       `json:"points" gorm:"default:0"`
	CreatorId int       `json:"creatorId"`
	Creator   User      `json:"creator"`
}

type PaginatedPosts struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"hasMore"`
}

type PostInput struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
