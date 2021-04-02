package db

import (
	"fmt"
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

func (p *PostTable) PostUpdate(id int, userId interface{}, options models.PostInput) (*models.Post, error) {
	var post models.Post

	p.Table.First(&post, id)
	if post.ID == 0 {
		return nil, fmt.Errorf("post not found")
	}

	userIdInt := userId.(int)
	if post.CreatorId != userIdInt {
		return nil, fmt.Errorf("you do not have permissions tp perform this action")
	}

	p.Table.Model(&post).Updates(&models.Post{
		Title: options.Title,
		Text:  options.Text,
	})

	return &post, nil
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
