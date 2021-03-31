package db

import (
	"lireddit/models"

	"gorm.io/gorm"
)

type PostTable struct {
	Table gorm.DB
}

func (p *PostTable) Postcreation(post models.Post) *models.Post {
	result := p.Table.Create(&post)
	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &post
}

func (p *PostTable) PostDelete(id int) bool {
	result := p.Table.Delete(&models.Post{}, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return false
	}

	return true
}

func (p *PostTable) PostUpdate(id int, title string) *models.Post {
	var post models.Post

	p.Table.First(&post, id)
	if post.ID == 0 {
		return nil
	}

	p.Table.Model(&post).Update("title", title)
	return &post
}

func (p *PostTable) GetPostById(id int) *models.Post {
	var post models.Post
	p.Table.First(&post, id)
	if post.ID == 0 {
		return nil
	}
	return &post
}

func (p *PostTable) GetAllPost() []models.Post {
	var posts []models.Post
	p.Table.Find(&posts)
	return posts
}
