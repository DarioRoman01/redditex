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

// delete post from the db and validate that the requesting user is owner of the post
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

// update post data and validate that requesting user is owner of the post
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
		p.Table.
			Table("posts").
			Where("created_at < ?", *cursor).
			Order("created_at DESC").
			Limit(limit).
			Preload("Creator").
			Find(&posts)
	} else {
		p.Table.
			Table("posts").
			Order("posts.created_at DESC").
			Limit(limit).
			Preload("Creator").
			Find(&posts)
	}

	if len(posts) == 0 {
		return nil, false
	}

	if len(posts) == limit {
		return posts[0 : limit-1], true
	}

	return posts[0 : len(posts)-1], false
}
