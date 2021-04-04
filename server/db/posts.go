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

func (p *PostTable) PostDelete(id int, userId int) error {
	var post models.Post
	p.Table.First(&post, id)

	if post.CreatorId != userId {
		return fmt.Errorf("you dont have permissions to perform this action")
	}

	if err := p.Table.Delete(&post); err != nil {
		return fmt.Errorf("unable to delete the post")
	}

	return nil
}

func (p *PostTable) PostUpdate(id int, userId int, options models.PostInput) (*models.Post, error) {
	var post models.Post

	p.Table.First(&post, id)
	if post.ID == 0 {
		return nil, fmt.Errorf("post not found")
	}

	if post.CreatorId != userId {
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

// get all posts and order the post by createdAt
// if a cursor is recibed only retrieves the posts created before that post
func (p *PostTable) GetAllPost(limit int, cursor *string) ([]models.Post, bool) {
	var posts []models.Post

	if limit > 50 {
		limit = 50
	}

	limit++
	if cursor != nil {
		p.Table.Where("created_at < ?", *cursor).Order("created_at DESC").Limit(limit).Find(&posts)
	} else {
		p.Table.Order("created_at DESC").Limit(limit).Find(&posts)
	}

	if len(posts) == limit {
		return posts, true
	}

	return posts, false
}
